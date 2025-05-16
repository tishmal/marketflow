package in

import (
	"context"
	"time"

	"marketflow/internal/domain/model"
)

// входной порт для API
type APIPort interface {
	// возвращает последнюю цену для указанной пары и биржи
	GetLatestPrice(ctx context.Context, pair string, exchange string) (model.PriceUpdate, error)
	
	// возвращает статистику цен для указанной пары и биржи за период
	GetPriceStats(ctx context.Context, pair string, exchange string, period time.Duration) (model.PriceStats, error)
	
	// возвращает список всех доступных бирж
	GetAllExchanges(ctx context.Context) ([]string, error)
	
	// возвращает список всех доступных торговых пар
	GetAllPairs(ctx context.Context) ([]string, error)
}