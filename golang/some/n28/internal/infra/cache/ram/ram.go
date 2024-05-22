package ram

import (
	"context"
	"errors"
	"n28/internal/domain"
	"sync"
	"time"
)

type cacheValue struct {
	value            any
	invalidationTime time.Time
}

type RamCache struct {
	mx     sync.Mutex
	values map[string]cacheValue
}

func New() (*RamCache, error) {
	return &RamCache{
		mx:     sync.Mutex{},
		values: make(map[string]cacheValue),
	}, nil
}

// public
// -----------------------------------------------------------------------

func (r *RamCache) SaveProduct(ctx context.Context,
	key string, product *domain.Product, ttl time.Duration) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.values[key] = cacheValue{
		value:            product,
		invalidationTime: time.Now().Add(ttl).UTC(),
	}

	return nil
}

func (r *RamCache) LoadProduct(ctx context.Context, key string) (*domain.Product, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	cv, exists := r.values[key]
	if !exists {
		return nil, errors.New("value by key not found")
	}
	if time.Now().UTC().Compare(cv.invalidationTime) != 1 {
		return nil, errors.New("value expired")
	}

	product, converted := cv.value.(*domain.Product)
	if !converted {
		return nil, errors.New("value is not a product")
	}

	return product, nil
}
