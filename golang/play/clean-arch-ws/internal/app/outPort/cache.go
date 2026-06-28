package outPort

import (
	"context"
	"golang/play/clean-arch-ws/internal/domain"
	"time"
)

type Cache interface {
	SaveProduct(ctx context.Context, key string, product *domain.Product, ttl time.Duration) error
	LoadProduct(ctx context.Context, key string) (*domain.Product, error)
}
