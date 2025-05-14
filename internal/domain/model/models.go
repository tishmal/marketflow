package model

import "time"

// метаданные о рынке
type MetaMarket struct {
	Exchange string // Имя биржи
	Pair     string // Торговая пара (например, "BTC/USD")
}

// обновление цен
type PriceUpdate struct {
	MetaMarket
	Price float64   // Текущая цена
	Time  time.Time // Время обновления
}

// статистика цен
type PriceStats struct {
	MetaMarket
	Timestamp time.Time // Время создания статистики
	Average   float64   // Средняя цена
	Min       float64   // Минимальная цена
	Max       float64   // Максимальная цена
}
