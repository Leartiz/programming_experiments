package useCase

import (
	"context"
	"errors"
	"fmt"
	"log"
	appDto "n28/internal/app/dto"
	"n28/internal/app/outPort"
	"n28/internal/domain"
	"time"
)

type Dependencies struct {
	Cache      outPort.Cache
	Db         outPort.Database
	Msgs       outPort.Msgs
	BkgTimeout time.Duration
}

func (d *Dependencies) validate() error {
	if d.Cache == nil {
		return errors.New("cache not provided")
	}
	if d.Db == nil {
		return errors.New("db not provided")
	}
	if d.Msgs == nil {
		return errors.New("msgs not provided")
	}

	return nil
}

type UseCase struct {
	cache      outPort.Cache
	db         outPort.Database
	msgs       outPort.Msgs
	bkgTimeout time.Duration

	productsChan chan *domain.Product
	ctxForCancel context.Context
}

func New(ctxForInit, ctxForCancel context.Context, deps Dependencies) (*UseCase, error) {
	if err := deps.validate(); err != nil {
		return nil, err
	}

	uc := UseCase{
		cache:        deps.Cache,
		db:           deps.Db,
		msgs:         deps.Msgs,
		bkgTimeout:   deps.BkgTimeout,
		productsChan: make(chan *domain.Product),
		ctxForCancel: ctxForCancel,
	}

	go uc.asyncConsumeExternalAddedProduct(ctxForInit, ctxForCancel)

	return &uc, nil
}

func (u *UseCase) InjectWs(ws outPort.Ws) error {
	return ws.SetChanForProductCountUpdated(u.ctxForCancel, u.productsChan)
}

// public
// -----------------------------------------------------------------------

func (u *UseCase) GetProductCount(ctx context.Context) (uint64, error) {
	return u.db.GetProductCount(ctx)
}

func (u *UseCase) InsertProductNoReturning(ctx context.Context, data appDto.InsertProduct) error {
	if data.Name == "" || data.Desc == "" || data.Price < 0 {
		return errors.New("invalid data")
	}

	// to infra...

	id, err := u.db.InsertProduct(ctx, appDto.InsertProduct{
		Name:  data.Name,
		Desc:  data.Desc,
		Price: data.Price,
	})
	if err != nil {
		return err
	}
	product, err := u.db.GetProductById(ctx, id)
	if err != nil {
		return err
	}

	err = u.cache.SaveProduct(ctx, fmt.Sprintf("product_%v", id), product, 30*time.Minute)
	if err != nil {
		return err
	}

	err = u.msgs.PublishNewProduct(ctx, "", product)
	if err != nil {
		return err
	}

	u.productsChan <- product

	return nil
}

// private
// -----------------------------------------------------------------------

func (u *UseCase) asyncConsumeExternalAddedProduct(ctxForInit, ctxForCancel context.Context) {
	ch, err := u.msgs.GetChanForProductAddedExternally(ctxForInit)
	if err != nil {
		log.Fatalf("failed to start asynchronous task")
	}

	for {
		select {
		case <-ctxForCancel.Done():
			log.Printf("async consume external added product [canceled]")
			return

		case product, ok := <-ch:
			if !ok {
				log.Printf("async consume external added product [closed]")
			}

			ctx, cancel := context.WithTimeout(context.Background(), u.bkgTimeout)
			defer cancel()

			u.InsertProductNoReturning(ctx, appDto.InsertProduct{
				Name:  product.Name,
				Desc:  product.Desc,
				Price: product.Price,
			})
		}
	}
}
