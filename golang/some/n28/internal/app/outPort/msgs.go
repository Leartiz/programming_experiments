package outPort

import (
	"context"
	"n28/internal/domain"
)

const (
	ExchangeName     = "n28"
	QueueProductName = "product"
	//...
)

type Msgs interface {
	PublishNewProduct(ctx context.Context, routingKey string, product *domain.Product) error
}
