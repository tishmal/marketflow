package out

import (
	"context"
	"time"

	"marketflow/internal/domain/model"
)

// порт для взаимодействия с хранилищем данных postrges
type StoragePort interface {
	// сохраняет информацию об обновлении цены
	SavePriceUpdate(ctx context.Context, update model.PriceUpdate) error

	// возвращает последнюю цену для указанной пары и биржи
	GetLatestPrice(ctx context.Context, pair, exchange string) (model.PriceUpdate, error)

	// возвращает историю цен для указанной пары и биржи за указанный период
	GetPriceHistory(ctx context.Context, pair, exchange string, from, to time.Time) ([]model.PriceUpdate, error)

	// возвращает список всех бирж в хранилище
	GetAllExchanges(ctx context.Context) ([]string, error)

	// возвращает список всех торговых пар в хранилище
	GetAllPairs(ctx context.Context) ([]string, error)
}
