package inPort

import (
	"context"
	inDto "n28/internal/app/inPort/dto"
)

type InsertProductChannel chan inDto.InsertProduct

func NewInsertProductChannel() chan inDto.InsertProduct {
	return make(chan inDto.InsertProduct)
}

type ProductUseCase interface {
	GetProductCount(ctx context.Context) (uint64, error)
	InsertProductNoReturning(ctx context.Context, data inDto.InsertProduct) error
}
