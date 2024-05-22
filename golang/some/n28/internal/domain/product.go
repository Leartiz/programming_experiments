package domain

import "errors"

type Product struct {
	Id    uint64
	Name  string
	Desc  string
	Price float64
}

func New(id uint64, name, desc string, price float64) (*Product, error) {

	// Domain Logic?

	if name == "" || desc == "" {
		return nil, errors.New("name or desc is empty")
	}
	if price < 0 {
		return nil, errors.New("negative price")
	}

	return &Product{
		Id:    id,
		Name:  name,
		Desc:  desc,
		Price: price,
	}, nil
}
