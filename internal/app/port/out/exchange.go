package port

import (
	"context"

	"marketflow/internal/domain/model"
)

// порт для взаимодействия с биржей
type ExchangePort interface {
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
	CreateExchange(name string) (ExchangePort, error)
	
	// возвращает список имен доступных бирж
	GetAvailableExchanges() []string
}