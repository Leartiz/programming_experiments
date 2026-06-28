package outPort

import (
	"context"
	"golang/play/clean-arch-ws/internal/domain"
)

type Ws interface {
	SetChanForProductCountUpdated(ctxForCancel context.Context,
		productsChan <-chan *domain.Product) error
}
