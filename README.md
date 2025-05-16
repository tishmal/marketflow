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
├── cmd/ # Точка входа
├── db/  # SQL скрипты
├── internal/
│ ├── domain/ # Модели и бизнес-логика
│ ├── application/ # Use cases и порты
│ ├── adapter/ # Входные и выходные адаптеры
│ └── config/ # Конфигурации
├── pkg/ # Вспомогательные пакеты (логгер, воркеры и т.д.)
├── deployments/ # Docker/Docker Compose
├── go.mod
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
docker-compose up --build
```

## 👨🏻‍💻 Автор

- [![Status](https://img.shields.io/badge/alem-tishmal-success?logo=github)](https://platform.alem.school/git/tishmal) <a href="https://t.me/tim_shm" target="_blank"><img src="https://img.shields.io/badge/telegram-@tishmal-blue?logo=Telegram" alt="Status" /></a>