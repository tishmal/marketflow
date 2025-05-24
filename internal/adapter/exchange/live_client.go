package exchange

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"marketflow/internal/domain/model"
	"net"
)

type LiveClient struct {
	ctx      context.Context
	addr     string
	exchange string
}

func NewLiveClient(ctx context.Context, addr, exchange string) *LiveClient {
	return &LiveClient{
		ctx:      ctx,
		addr:     addr,
		exchange: exchange,
	}
}

func (e *LiveClient) Subscribe(ctx context.Context, pairs []string) (<-chan model.PriceUpdate, <-chan error, error) {
	prices := make(chan model.PriceUpdate)
	errs := make(chan error)

	conn, err := net.Dial("tcp", e.addr)
	if err != nil {
		return nil, nil, err
	}

	go func() {
		defer conn.Close()
		defer close(prices)
		defer close(errs)

		scanner := bufio.NewScanner(conn)

		for {
			select {
			case <-ctx.Done():
				return
			default:
				if scanner.Scan() {
					line := scanner.Text()
					var update model.PriceUpdate

					err := json.Unmarshal([]byte(line), &update)
					if err != nil {
						errs <- errors.New("failed to parse price update: " + err.Error())
						continue
					}

					// фильтрация по нужным парам
					for _, p := range pairs {
						if update.Pair == p {
							prices <- update
							break
						}
					}
				} else if err := scanner.Err(); err != nil {
					errs <- err
					return
				}
			}
		}
	}()

	return prices, errs, nil
}

func (e *LiveClient) Close() error {
	// по факту закрывается внутри goroutine
	return nil
}

func (e *LiveClient) GetName() string {
	return e.exchange
}
