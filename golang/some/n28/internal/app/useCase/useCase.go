package useCase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"n28/internal/app/inPort"
	inDto "n28/internal/app/inPort/dto"
	"n28/internal/app/outPort"
	outDto "n28/internal/app/outPort/dto"
	"time"
)

type Dependencies struct {
	Db    outPort.Database
	Msgs  outPort.Msgs
	Cache outPort.Cache

	InsertProductChannel inPort.InsertProductChannel
}

func (d *Dependencies) validate() error {
	if d.Db == nil {
		return errors.New("db not provided")
	}
	if d.Msgs == nil {
		return errors.New("msgs not provided")
	}
	if d.Cache == nil {
		return errors.New("cache not provided")
	}

	if d.InsertProductChannel == nil {
		return errors.New("`insert product` channel not provided")
	}

	return nil
}

type UseCase struct {
	db              outPort.Database
	msgs            outPort.Msgs
	cache           outPort.Cache
	insertProductCh inPort.InsertProductChannel
}

func New(deps Dependencies) (*UseCase, error) {
	if err := deps.validate(); err != nil {
		return nil, err
	}

	return &UseCase{
		db:    deps.Db,
		msgs:  deps.Msgs,
		cache: deps.Cache,
	}, nil
}

// public
// -----------------------------------------------------------------------

func (u *UseCase) GetProductCount(ctx context.Context) (uint64, error) {
	return u.db.GetProductCount(ctx)
}

func (u *UseCase) InsertProductNoReturning(ctx context.Context, data inDto.InsertProduct) error {
	if data.Name == "" || data.Desc == "" || data.Price < 0 {
		return errors.New("invalid data")
	}

	// ***

	id, err := u.db.InsertProduct(ctx, outDto.InsertProduct{
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

	// ***

	err = u.cache.SaveProduct(ctx, fmt.Sprintf("product_%v", id), product, 30*time.Minute)
	if err != nil {
		return err
	}

	err = u.msgs.PublishNewProduct(ctx, "", product)
	if err != nil {
		return err
	}

	return nil
}

// private
// -----------------------------------------------------------------------

func (u *UseCase) bkgProcessingInsertProduct() {
	for {
		select {
		case data, closed := <-u.insertProductCh:
			if closed {
				return
			}

			err := u.InsertProductNoReturning(context.Background(), data)
			if err != nil {
				log.Printf("bkg insert product err %v", err)
			}
		}
	}
}
