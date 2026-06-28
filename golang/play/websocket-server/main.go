package main

import (
	"log"
	"net/http"
	"time"

	"github.com/lxzan/gws"
)

const (
	PingTimeout = 5 * time.Second
)

type FuncWithReturningError func() error

func RunFnsErrs(fns ...FuncWithReturningError) error {
	for i := range fns {
		if err := fns[i](); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	upgrader := gws.NewUpgrader(&Handler{}, &gws.ServerOption{
		ParallelEnabled:   true,                                 // Parallel message processing
		Recovery:          gws.Recovery,                         // Exception recovery
		PermessageDeflate: gws.PermessageDeflate{Enabled: true}, // Enable compression
	})
	http.HandleFunc("/connect", func(writer http.ResponseWriter, request *http.Request) {
		socket, err := upgrader.Upgrade(writer, request)
		if err != nil {
			return
		}
		go func() {
			socket.ReadLoop() // Blocking prevents the context from being GC.
		}()
	})
	http.ListenAndServe(":6666", nil)
}

// -----------------------------------------------------------------------

type Handler struct{}

func (c *Handler) OnOpen(socket *gws.Conn) {
	log.Println("on open")
	go c.pingRoutine(socket)
}

func (c *Handler) OnClose(socket *gws.Conn, err error) {
	log.Printf("on close: %v\n", err)
}

func (c *Handler) OnPing(socket *gws.Conn, payload []byte) {
	log.Println("on ping")

	err := RunFnsErrs(
		func() error {
			return socket.SetDeadline(time.Now().Add(PingTimeout))
		},
		func() error {
			log.Println("send pong")
			return socket.WritePong(nil)
		})

	if err != nil {
		log.Printf("%v", err)
	}
}

func (c *Handler) OnPong(socket *gws.Conn, payload []byte) {
	log.Printf("on pong: '%v'\n", string(payload))
}

func (c *Handler) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()

	log.Printf("on message: '%v'\n", string(message.Bytes()))

	err := RunFnsErrs(
		func() error {
			return socket.SetWriteDeadline(time.Now().Add(PingTimeout))
		},
		func() error {
			log.Println("send msg")
			return socket.WriteMessage(message.Opcode, message.Bytes())
		})

	if err != nil {
		log.Printf("%v", err)
	}
}

func (c *Handler) pingRoutine(socket *gws.Conn) {
	for {
		time.Sleep(2500 * time.Millisecond)
		err := RunFnsErrs(
			func() error {
				return socket.SetDeadline(time.Now().Add(PingTimeout))
			},
			func() error {
				return socket.WritePing(nil)
			},
		)

		if err != nil {
			log.Printf("ping routine error: %v", err)
			return
		}
	}
}
