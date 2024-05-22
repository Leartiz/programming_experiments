package rabbit

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"n28/internal/app/inPort"
	inPortDto "n28/internal/app/inPort/dto"
	"n28/internal/domain"
	rabbitDto "n28/internal/infra/msgs/rabbit/dto"

	"github.com/wagslane/go-rabbitmq"
)

type Dependencies struct {
	User string
	Pass string
	Host string
	Port uint16

	ExchangeName                   string
	QueueNameExternalAddedProducts string

	ChInsertProductChannel inPort.InsertProductChannel
}

func createConnString(deps Dependencies) string {
	return fmt.Sprintf("amqp://%v:%v@%v:%v",
		deps.User, deps.Pass, deps.Host, deps.Port)
}

type RabbitMsgs struct {
	conn      *rabbitmq.Conn
	publisher *rabbitmq.Publisher
	consumer  *rabbitmq.Consumer

	exchangeName                   string
	queueNameExternalAddedProducts string

	chExternalAddedProducts chan string
	chInsertProductChannel  inPort.InsertProductChannel
}

func New(deps Dependencies) (*RabbitMsgs, error) {
	connString := createConnString(deps)
	fmt.Printf("conn string %v", connString)

	conn, err := rabbitmq.NewConn(
		connString,
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		return nil, err
	}

	// -------------------------------------------------------------------

	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeName(deps.ExchangeName),
		rabbitmq.WithPublisherOptionsExchangeDeclare,
	)
	if err != nil {
		return nil, err
	}

	// -------------------------------------------------------------------

	consumer, err := rabbitmq.NewConsumer(
		conn,
		deps.QueueNameExternalAddedProducts,
		rabbitmq.WithConsumerOptionsRoutingKey(""),
		rabbitmq.WithConsumerOptionsExchangeName(deps.ExchangeName),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
	)
	if err != nil {
		return nil, err
	}

	chExternalAddedProducts := make(chan string)
	err = consumer.Run(func(d rabbitmq.Delivery) rabbitmq.Action {
		chExternalAddedProducts <- string(d.Body)
		return rabbitmq.Ack
	})
	if err != nil {
		return nil, err
	}

	// -------------------------------------------------------------------

	instance := RabbitMsgs{
		conn:      conn,
		publisher: publisher,
		consumer:  consumer,

		exchangeName:                   deps.ExchangeName,
		queueNameExternalAddedProducts: deps.QueueNameExternalAddedProducts,

		chExternalAddedProducts: chExternalAddedProducts,
		chInsertProductChannel:  deps.ChInsertProductChannel,
	}

	//go instance.consumeExternalAddedProducts()
	return &instance, nil
}

// public
// -----------------------------------------------------------------------

func (r *RabbitMsgs) PublishNewProduct(ctx context.Context,
	routingKey string, product *domain.Product) error {

	jsonBytes, err := json.Marshal(product)
	if err != nil {
		return err
	}

	err = r.publisher.Publish(
		jsonBytes, []string{routingKey},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsExchange(r.exchangeName),
	)
	if err != nil {
		return err
	}
	log.Printf("product published")

	return nil
}

// private
// -----------------------------------------------------------------------

func (r *RabbitMsgs) consumeExternalAddedProducts() {
	for {
		select {
		case msg, closed := <-r.chExternalAddedProducts:
			if closed {
				return
			}

			product := rabbitDto.AddNewProduct{}
			err := json.Unmarshal([]byte(msg), &product)
			if err != nil {
				log.Printf("failed to consume product")
				continue
			}

			r.chInsertProductChannel <- inPortDto.InsertProduct{
				Name:  product.Name,
				Desc:  product.Desc,
				Price: product.Price,
			}
		}
	}
}
