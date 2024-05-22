package launch

import (
	"log"
	"n28/internal/app/inPort"
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

func Launch() error {

	// Async In Ports To Infra

	insertProductCh := inPort.NewInsertProductChannel()

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
	})
	if err != nil {
		return nil
	}
	log.Printf("msgs [OK]")

	// App

	ucDeps := useCase.Dependencies{
		Db:    outPortDb,
		Cache: outPortCache,
		Msgs:  outPortMsgs,

		InsertProductChannel: insertProductCh,
	}

	uc, err := useCase.New(ucDeps)
	if err != nil {
		return nil
	}

	// Deps For Infra

	// provide access to the client at the very end!!

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
