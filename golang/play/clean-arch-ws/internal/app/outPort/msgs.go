package outPort

import (
	"context"
	"golang/play/clean-arch-ws/internal/app/dto"
	"golang/play/clean-arch-ws/internal/domain"
)

const (
	ExchangeName                  = "n28"
	QueueNameExternalAddedProduct = "product"
)

type Msgs interface {
	PublishNewProduct(ctx context.Context, routingKey string, product *domain.Product) error
	GetChanForProductAddedExternally(ctx context.Context) (<-chan dto.InsertProduct, error)
}
