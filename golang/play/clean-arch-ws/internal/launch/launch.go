package launch

import (
	"context"
	"log"
	"golang/play/clean-arch-ws/internal/app/outPort"
	"golang/play/clean-arch-ws/internal/app/useCase"
	"golang/play/clean-arch-ws/internal/infra/cache/ram"
	"golang/play/clean-arch-ws/internal/infra/msgs/rabbit"
	"golang/play/clean-arch-ws/internal/infra/sql/lite"
	"golang/play/clean-arch-ws/internal/infra/ws"
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

	wsHandler, err := ws.NewHandler(ws.Dependencies{
		PingInterval: PingInterval,
		PingWait:     PingWait,
		ProductUc:    uc,
	})
	if err != nil {
		return err
	}

	uc.InjectWs(wsHandler)

	// Deps For Infra

	err = ws.Listen(wsHandler)
	if err != nil {
		return err
	}

	return nil
}
