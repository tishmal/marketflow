package domain

import "time"

// обновление цен
type PriceUpdate struct {
	Exchange string    // Имя биржи
	Pair     string    // Торговая пара (например, "BTC/USD")
	Price    float64   // Текущая цена
	Time     time.Time // Время обновления
}

// статистика цен
type PriceStats struct {
	Exchange  string
	Pair      string
	Timestamp time.Time // Время создания статистики
	Average   float64   // Средняя цена
	Min       float64   // Минимальная цена
	Max       float64   // Максимальная цена
}
