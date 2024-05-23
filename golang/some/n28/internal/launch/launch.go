package launch

import (
	"context"
	"log"
	"n28/internal/app/outPort"
	"n28/internal/app/useCase"
	"n28/internal/infra/cache/ram"
	"n28/internal/infra/msgs/rabbit"
	"n28/internal/infra/sql/lite"
	"n28/internal/infra/ws"
	"time"
)

// From Config!

const (
	PingInterval = 5 * time.Second
	PingWait     = 10 * time.Second
)

func Launch(ctxForInit, ctxForCancel context.Context) error {

	// Infra Out Ports

	outPortDb, err := lite.New()
	if err != nil {
		return nil
	}
	log.Printf("db [OK]")

	outPortCache, err := ram.New()
	if err != nil {
		return nil
	}
	log.Printf("cache [OK]")

	outPortMsgs, err := rabbit.New(rabbit.Dependencies{
		User: "guest",
		Pass: "guest",
		Host: "localhost",
		Port: 5672,

		ExchangeName:                   outPort.ExchangeName,
		QueueNameExternalAddedProducts: outPort.QueueNameExternalAddedProduct,
	})
	if err != nil {
		return nil
	}
	log.Printf("msgs [OK]")

	// App

	ucDeps := useCase.Dependencies{
		Db:         outPortDb,
		Cache:      outPortCache,
		Msgs:       outPortMsgs,
		BkgTimeout: 5 * time.Second,
	}

	uc, err := useCase.New(ctxForInit, ctxForCancel, ucDeps)
	if err != nil {
		return nil
	}

	// Deps For Infra

	err = ws.Listen(ws.Dependencies{
		PingInterval: PingInterval,
		PingWait:     PingWait,
		ProductUc:    uc,
	})
	if err != nil {
		return err
	}

	return nil
}
