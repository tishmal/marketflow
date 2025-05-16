package out

import (
	"context"

	"marketflow/internal/domain/model"
)

// порт для взаимодействия с кэшем
type CachePort interface {
	// кэширует последнюю цену
	SetLatestPrice(ctx context.Context, update model.PriceUpdate) error
	
	// получает последнюю цену из кэша
	GetLatestPrice(ctx context.Context, pair, exchange string) (model.PriceUpdate, bool, error)
	
	// кэширует статистику цен
	SetPriceStats(ctx context.Context, stats model.PriceStats) error
	
	// получает статистику цен из кэша
	GetPriceStats(ctx context.Context, pair, exchange string) (model.PriceStats, bool, error)
}