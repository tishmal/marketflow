package out

import (
	"context"

	"marketflow/internal/domain/model"
)

type ExchangeClient interface {
	// устанавливает соединение с биржей
	Connect(ctx context.Context) error

	// закрывает соединение с биржей
	Close() error

	// подписывается на обновления цен и возвращает канал с обновлениями
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
