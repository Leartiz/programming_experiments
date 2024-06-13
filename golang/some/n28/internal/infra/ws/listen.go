package ws

import (
	"log"
	"n28/internal/app/inPort"
	"net/http"
	"time"

	"github.com/lxzan/gws"
)

type Dependencies struct {
	PingInterval time.Duration
	PingWait     time.Duration
	ProductUc    inPort.ProductUseCase
}

func Listen(handler *Handler) error {
	upgrader := gws.NewUpgrader(handler, &gws.ServerOption{
		ParallelEnabled: true,
		Recovery:        gws.Recovery,
	})

	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		socket, err := upgrader.Upgrade(writer, request)
		if err != nil {
			return
		}

		go func() {
			socket.ReadLoop()
		}()
	})

	const addr = ":6666"
	go func() {
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatalf("failed to start ws listener")
		}
	}()

	log.Printf("ws will be listening on %v", addr)
	return nil
}
