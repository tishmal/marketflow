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

func (m *Manager) Start(ctx context.Context, outData chan<- model.PriceUpdate, mode Mode) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.cancelFunc != nil {
		logger.Info("stopping previous mode", "mode", m.mode)
		m.cancelFunc()
		for _, client := range m.clients {
			client.Close()
		}
	}

	ctx, cancel := context.WithCancel(ctx)
	m.cancelFunc = cancel
	m.mode = mode
	m.clients = nil

	switch mode {
	case Test:
		// Здесь можно добавить генераторы фейковых данных, если нужно
		// m.clients = append(m.clients, exchange.NewTestGenerator("ex1"))
	case Live:
		for _, ex := range m.cfg.Exchanges {
			client := exchange.NewLiveClient(ctx, ex.Addr, ex.Name)
			m.clients = append(m.clients, client)
		}
	default:
		return errors.New("invalid mode")
	}
	// заглушка:
	pairs := []string{"BTCUSDT"}

	for _, client := range m.clients {
		go func(c out.ExchangeClient) {
			prices, errs, err := c.Subscribe(ctx, pairs)
			if err != nil {
				logger.Error("failed to subscribe to client", "client", c.GetName(), "error", err)
				return
			}

			for {
				select {
				case price, ok := <-prices:
					if !ok {
						return
					}
					outData <- price
				case err, ok := <-errs:
					if ok {
						logger.Error("error from exchange", "client", c.GetName(), "error", err)
					}
				case <-ctx.Done():
					return
				}
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
			client.Close()
		}
		m.cancelFunc = nil
		m.clients = nil
	}
}

func (m *Manager) Current() Mode {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.mode
}
