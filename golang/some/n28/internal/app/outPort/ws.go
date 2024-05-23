package outPort

import (
	"context"
	"n28/internal/domain"
)

type Ws interface {
	SetChanForProductCountUpdated(ctxForCancel context.Context,
		productsChan chan<- *domain.Product) error
}
