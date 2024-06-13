package outPort

import (
	"context"
	appDto "n28/internal/app/dto"
	"n28/internal/domain"
)

type Database interface {
	InsertProduct(ctx context.Context, data appDto.InsertProduct) (uint64, error)
	GetProductById(ctx context.Context, id uint64) (*domain.Product, error)
	GetProductCount(ctx context.Context) (uint64, error)
}
