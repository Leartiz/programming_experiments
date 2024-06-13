package ws

import (
	"context"
	"encoding/json"
	"log"
	appDto "n28/internal/app/dto"
	"n28/internal/app/inPort"
	"n28/internal/domain"
	wsDto "n28/internal/infra/ws/dto"
	"time"

	"github.com/lxzan/gws"
)

type Handler struct {
	pingInterval time.Duration
	pingWait     time.Duration
	productUc    inPort.ProductUseCase
	sockets      []*gws.Conn
}

func NewHandler(deps Dependencies) (*Handler, error) {
	h := Handler{
		pingInterval: deps.PingInterval,
		pingWait:     deps.PingWait,
		productUc:    deps.ProductUc,
	}
	return &h, nil
}

// public
// -----------------------------------------------------------------------

func (c *Handler) OnOpen(socket *gws.Conn) {
	_ = socket.SetDeadline(time.Now().Add(c.pingInterval + c.pingWait))
	c.sockets = append(c.sockets, socket)
}

func (c *Handler) OnClose(socket *gws.Conn, err error) {
	log.Printf("conn %v close", socket.RemoteAddr())

	var index int
	for i := range c.sockets {
		if c.sockets[i] == socket {
			index = i
			continue
		}
	}

	c.sockets = append(c.sockets[:index],
		c.sockets[index+1:]...)
}

/// From Library
// -----------------------------------------------------------------------

func (c *Handler) OnPing(socket *gws.Conn, payload []byte) {
	_ = socket.SetDeadline(time.Now().Add(c.pingInterval + c.pingWait))
	_ = socket.WritePong(nil)
}

func (c *Handler) OnPong(socket *gws.Conn, payload []byte) {}

func (c *Handler) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()

	messageStruct := wsDto.Message{}
	err := json.Unmarshal(message.Bytes(), &messageStruct)
	if err != nil {
		socket.WriteString("400 (json)")
		return
	}

	// ***

	if messageStruct.Event == "add_new_product" {
		productBytes, err := json.Marshal(messageStruct.Payload)
		if err != nil {
			socket.WriteString("400 (json)")
			return
		}

		product := wsDto.AddNewProduct{}
		err = json.Unmarshal(productBytes, &product)
		if err != nil {
			socket.WriteString("400 (json)")
			return
		}

		err = c.productUc.InsertProductNoReturning(context.Background(), appDto.InsertProduct{
			Name:  product.Name,
			Desc:  product.Desc,
			Price: product.Price,
		})
		if err != nil {
			socket.WriteString("450 (server or client)")
			return
		}

		socket.WriteString("200")
		return
	}

	// Echo Mode...

	msgText := string(message.Bytes())
	log.Printf("received msg %v from %v", msgText, socket.RemoteAddr())

	msgText = "*** BEGIN ***\n" + msgText + "\n*** END ***"
	socket.WriteString(msgText)
}

// use case
// -----------------------------------------------------------------------

func (c *Handler) SetChanForProductCountUpdated(ctxForCancel context.Context,
	productsChan <-chan *domain.Product) error {
	go func() {
		for {
			select {
			case <-ctxForCancel.Done():
				log.Printf("bkg work canceled [product count updated]")
				return

			case product, ok := <-productsChan:
				if !ok {
					log.Printf("bkg work canceled [product count updated]")
					return
				}

				jsonProduct, err := json.Marshal(product)
				if err != nil {
					log.Printf("")
					continue
				}

				for i := range c.sockets {
					c.sockets[i].WriteString(string(jsonProduct))
				}
			}
		}
	}()

	return nil
}
