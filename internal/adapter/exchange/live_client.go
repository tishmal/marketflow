package exchange

import (
	"context"
	"marketflow/internal/domain/model"
	"time"
)

type LiveAdapter struct {
	ctx      context.Context
	addr     string
	exchange string
}

func NewLiveClient(ctx context.Context, addr, exchange string) *LiveAdapter {
	return &LiveAdapter{
		ctx:      ctx,
		addr:     addr,
		exchange: exchange,
	}
}

func (e *LiveAdapter) Connect(ctx context.Context) error {
	// logic
	return nil
}

func (e *LiveAdapter) Subscribe(ctx context.Context, pairs []string) (<-chan model.PriceUpdate, <-chan error, error) {
	prices := make(chan model.PriceUpdate)
	errs := make(chan error)

	go func() {
		ticker := time.NewTicker(3 * time.Second)
		for {
			select {
			case <-ctx.Done():
				close(prices)
				close(errs)
				return
			case t := <-ticker.C:
				prices <- model.PriceUpdate{
					// ...
				}
			}
		}
	}()

	return prices, errs, nil
}

func (e *LiveAdapter) Close() error {
	// ничего не делаем
	return nil
}

func (e *LiveAdapter) GetName() string {
	return e.exchange
}
