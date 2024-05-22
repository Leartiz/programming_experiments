package outPort

import (
	"context"
	"n28/internal/domain"
	"time"
)

type Cache interface {
	SaveProduct(ctx context.Context, key string, product *domain.Product, ttl time.Duration) error
	LoadProduct(ctx context.Context, key string) (*domain.Product, error)
}
