package out

import (
	"context"

	"marketflow/internal/domain/model"
)

type ExchangeClient interface {
	// закрывает соединение с биржей
	Close() error

	// устанавливает соединение с биржей и подписывается на обновления цен и возвращает канал с обновлениями
	Subscribe(ctx context.Context, pairs []string) (<-chan model.PriceUpdate, <-chan error, error)

	// возвращает имя биржи
	GetName() string
}

// создает экземпляры бирж
type ExchangeFactory interface {
	// создает адаптер для биржи
	CreateExchange(name string) (ExchangeClient, error)

	// возвращает список имен доступных бирж
	GetAvailableExchanges() []string
}
