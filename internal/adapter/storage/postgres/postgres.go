package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"marketflow/internal/config"
	"marketflow/internal/domain/model"
	"marketflow/pkg/logger"
	"time"

	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(cfg config.PostgresConfig) (*PostgresRepository, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.NameDB, cfg.SSLMode,
	)
	// ленивое соединение
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error("failed to connect to postgres", "error", err)
		return nil, fmt.Errorf("cannot open db: %w", err)
	}
	// проверяет действительно ли можно подключиться и реальное соединение происходит
	if err := db.Ping(); err != nil {
		logger.Error("failed to ping postgres", "error", err)
		return nil, fmt.Errorf("cannot connect to db: %w", err)
	}
	// ограничения
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Hour)

	logger.Info("postgres connection success")
	return &PostgresRepository{db: db}, nil
}

func (r *PostgresRepository) SavePriceUpdate(ctx context.Context, update model.PriceUpdate) error {
	return nil
}

func (r *PostgresRepository) Close() error {
	logger.Info("postgres db is close")
	return r.db.Close()
}
