package inPort

import (
	"context"
	appDto "golang/play/clean-arch-ws/internal/app/dto"
)

type ProductUseCase interface {
	GetProductCount(ctx context.Context) (uint64, error)
	InsertProductNoReturning(ctx context.Context, data appDto.InsertProduct) error
}
