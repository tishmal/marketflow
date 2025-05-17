# Marketflow 📁

![Go](https://img.shields.io/badge/Go-1.23-violet) ![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-violet) ![Docker](https://img.shields.io/badge/Docker-✓-violet)

**Marketflow** это высоконагруженная система для обработки рыночных данных в реальном времени. Она реализована по принципам **гексагональной архитектуры (Hexagonal Architecture / Ports and Adapters)**, обеспечивая чистоту кода, расширяемость и легкость тестирования.

## ✨ Возможности

- Получение актуальных рыночных цен по валютным парам
- Агрегация статистики (минимум, максимум, среднее)
- Поддержка live/test режимов
- Кэширование через Redis
- Сохранение данных в PostgreSQL
- Гибкое подключение к различным источникам биржевых данных

## 📚 Технологии

- **Go** 1.23+
- **PostgreSQL**
- **Redis**
- **REST API**
- **Docker**

## 🧬 Структура проекта

marketflow/
├── cmd/
│   └── marketflow/
│       └── main.go              # Точка входа в приложение
├── db/
│   ├── migrations/              # SQL-скрипты для миграций
│   └── init.sql                 # Начальная инициализация БД
├── internal/
│   ├── domain/                  # Модели предметной области и бизнес-логика
│   │   ├── model/               # Доменные модели
│   │   └── service/             # Бизнес-логика
│   ├── application/             # Use cases, интерфейсы портов
│   │   ├── port/
│   │   │   ├── input/           # Входные порты (интерфейсы для API)
│   │   │   └── output/          # Выходные порты (интерфейсы для хранилищ)
│   │   └── usecase/             # Реализация сценариев использования
│   ├── adapter/                 # Входные/выходные адаптеры
│   │   ├── handler/             # REST API обработчики
│   │   ├── repository/          # Репозитории для работы с БД
│   │   ├── exchange/            # Адаптеры для биржевых данных
│   │   └── cache/               # Адаптер для работы с Redis
│   └── config/                  # Загрузка и хранение конфигураций
├── pkg/
│   ├── logger/                  # Логгирование с использованием log/slog
│   ├── worker/                  # Реализация Worker Pool
│   ├── concurrency/             # Паттерны конкурентного программирования
│   └── utils/                   # Вспомогательные функции
├── deployments/
│   ├── docker/                  # Dockerfile для приложения
│   └── docker-compose.yml       # Конфигурация Docker Compose
├── test/
│   ├── generator/               # Генератор тестовых данных
│   └── integration/             # Интеграционные тесты
├── .gitignore
├── go.mod
├── go.sum
├── Makefile
└── README.md

## 🧱 Архитектура

Проект построен по принципу **Hexagonal Architecture**:
- **Domain Layer**: бизнес-сущности и их поведение
- **Application Layer**: use cases и порты (контракты взаимодействия)
- **Adapters Layer**:
  - `in`: HTTP API, биржевые коннекторы
  - `out`: PostgreSQL, Redis, внешние биржи


## 🏁 Запуск (локально)

```bash
make up
```

## 👨🏻‍💻 Автор

- [![Status](https://img.shields.io/badge/alem-tishmal-success?logo=github)](https://platform.alem.school/git/tishmal) <a href="https://t.me/tim_shm" target="_blank"><img src="https://img.shields.io/badge/telegram-@tishmal-blue?logo=Telegram" alt="Status" /></a>