package ws

import (
	"context"
	"encoding/json"
	"log"
	"n28/internal/app/inPort"
	inPortDto "n28/internal/app/inPort/dto"
	wsDto "n28/internal/infra/ws/dto"
	"time"

	"github.com/lxzan/gws"
)

type Handler struct {
	pingInterval time.Duration
	pingWait     time.Duration
	productUc    inPort.ProductUseCase
}

func NewHandler(deps Dependencies) (*Handler, error) {
	return &Handler{
		pingInterval: deps.PingInterval,
		pingWait:     deps.PingWait,
		productUc:    deps.ProductUc,
	}, nil
}

// public
// -----------------------------------------------------------------------

func (c *Handler) OnOpen(socket *gws.Conn) {
	_ = socket.SetDeadline(time.Now().Add(c.pingInterval + c.pingWait))
}

func (c *Handler) OnClose(socket *gws.Conn, err error) {
	log.Printf("conn %v close", socket.RemoteAddr())
}

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

		err = c.productUc.InsertProductNoReturning(context.Background(), inPortDto.InsertProduct{
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

// private
// -----------------------------------------------------------------------
