package inPort

import (
	"context"
	inDto "n28/internal/app/inPort/dto"
)

type ProductUseCase interface {
	GetProductCount(ctx context.Context) (uint64, error)
	InsertProductNoReturning(ctx context.Context, data inDto.InsertProduct) error
}
