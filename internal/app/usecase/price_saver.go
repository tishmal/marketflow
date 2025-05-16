package usecase

import (
	"context"
	"marketflow/internal/app/port/out"
	"marketflow/internal/domain/model"
)

type PriceSaver struct {
	storage out.StoragePort
}

func NewPriceSaver(storage out.StoragePort) *PriceSaver {
	return &PriceSaver{storage: storage}
}

func (ps *PriceSaver) Save(ctx context.Context, updatePrice model.PriceUpdate) error {
	return ps.storage.SavePriceUpdate(ctx, updatePrice)
}

