package outPort

import (
	"context"
	"n28/internal/app/dto"
	"n28/internal/domain"
)

const (
	ExchangeName                  = "n28"
	QueueNameExternalAddedProduct = "product"
)

type Msgs interface {
	PublishNewProduct(ctx context.Context, routingKey string, product *domain.Product) error
	GetChanForProductAddedExternally(ctx context.Context) (<-chan dto.InsertProduct, error)
}
