package postgres

import (
	"context"
	"database/sql"
	"marketflow/internal/domain/model"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}

func (r *Repository) SavePriceUpdate(ctx context.Context, update model.PriceUpdate) error {
}