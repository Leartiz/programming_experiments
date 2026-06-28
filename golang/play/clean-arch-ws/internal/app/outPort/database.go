package outPort

import (
	"context"
	appDto "golang/play/clean-arch-ws/internal/app/dto"
	"golang/play/clean-arch-ws/internal/domain"
)

type Database interface {
	InsertProduct(ctx context.Context, data appDto.InsertProduct) (uint64, error)
	GetProductById(ctx context.Context, id uint64) (*domain.Product, error)
	GetProductCount(ctx context.Context) (uint64, error)
}
