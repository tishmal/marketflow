package out

import (
	"context"
)

type CachePort interface {
	// Connect устанавливает соединение с кешем
	ConnectCache(ctx context.Context) error
	// Close закрывает соединение с кешем
	Close() error
	// // SavePriceUpdate сохраняет обновление цены в кеше
	// SavePriceUpdate(ctx context.Context, update model.PriceUpdate) error
	// // GetLatestPrice получает последнюю цену для указанной пары
	// GetLatestPrice(ctx context.Context, pair string) (model.PriceResponse, error)
	// // GetLatestPriceByExchange получает последнюю цену для указанной пары и биржи
	// GetLatestPriceByExchange(ctx context.Context, exchange, pair string) (model.PriceResponse, error)
	// // GetPriceUpdatesInRange получает обновления цен за указанный период времени
	// GetPriceUpdatesInRange(ctx context.Context, pair string, duration time.Duration) ([]model.PriceUpdate, error)
	// // GetPriceUpdatesInRangeByExchange получает обновления цен за указанный период времени для указанной биржи
	// GetPriceUpdatesInRangeByExchange(ctx context.Context, exchange, pair string, duration time.Duration) ([]model.PriceUpdate, error)
	// // CleanupOldData удаляет устаревшие данные из кеша
	// CleanupOldData(ctx context.Context) error
	// // IsHealthy проверяет работоспособность кеша
	// IsHealthy(ctx context.Context) bool
}
