package mode

import (
	"context"
	"errors"
	"sync"

	"marketflow/internal/adapter/exchange"
	"marketflow/internal/app/port/out"
	"marketflow/internal/config"
	"marketflow/internal/domain/model"
	"marketflow/pkg/logger"
)

type Mode string

const (
	Live Mode = "live"
	Test Mode = "test"
)

type Manager struct {
	mu         sync.Mutex
	mode       Mode
	clients    []out.ExchangeClient
	cancelFunc context.CancelFunc
	cfg        *config.Config
}

func NewManager(cfg *config.Config) *Manager {
	return &Manager{
		mode: Test,
		cfg:  cfg,
	}
}

func (m *Manager) Start(ctx context.Context, out chan<- model.PriceUpdate, mode Mode) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.cancelFunc != nil {
		logger.Info("stopping previous mode")
		m.cancelFunc()
		for _, client := range m.clients {
			client.Close()
		}
	}

	ctx, cancel := context.WithCancel(ctx)
	m.cancelFunc = cancel
	m.mode = mode

	m.clients = nil
	if mode == Test {
		m.clients = []out.ExchangeClient{
			exchange.NewLiveClient()

			// exchange.NewTestGenerator("ex1"),
			// exchange.NewTestGenerator("ex2"),
			// exchange.NewTestGenerator("ex3"),
		}
	} else if mode == Live {
		for _, ex := range m.cfg.Exchanges {
			m.clients = append(m.clients, exchange.NewTCPClient(ctx, ex.Name, ex.Address))
		}
	} else {
		return errors.New("invalid mode")
	}

	for _, client := range m.clients {
		go func(c domain.ExchangeClient) {
			if err := c.Start(ctx, out); err != nil {
				logger.Error("failed to start client", "client", c, "error", err)
			}
		}(client)
	}

	logger.Info("started mode", "mode", mode)
	return nil
}

func (m *Manager) Stop() {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.cancelFunc != nil {
		logger.Info("stopping mode", "mode", m.mode)
		m.cancelFunc()
		for _, client := range m.clients {
			client.Stop()
		}
		m.cancelFunc = nil
		m.clients = nil
	}
}

func (m *Manager) Current() Mode {
	m.mu.Lock()
	defer m.mu.Unlock()
	logger.Info("current mode", "mode", m.mode)
	return m.mode
}
