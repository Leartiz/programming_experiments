package rabbit

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	appDto "n28/internal/app/dto"
	"n28/internal/domain"
	rabbitDto "n28/internal/infra/msgs/rabbit/dto"
	"time"

	"github.com/wagslane/go-rabbitmq"
)

type Dependencies struct {
	User string
	Pass string
	Host string
	Port uint16

	ExchangeName                   string
	QueueNameExternalAddedProducts string
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

	chanForProductAddedExternally chan appDto.InsertProduct
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
		rabbitmq.WithConsumerOptionsExchangeName(deps.ExchangeName),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
	)
	if err != nil {
		return nil, err
	}

	externalAddedProductsChan := make(chan appDto.InsertProduct, 5)
	go func() {
		err := consumer.Run(func(d rabbitmq.Delivery) rabbitmq.Action {
			product := rabbitDto.AddNewProduct{}
			err := json.Unmarshal(d.Body, &product)
			if err != nil {
				log.Printf("failed to consume product")
				return rabbitmq.NackDiscard
			}

			select {
			case externalAddedProductsChan <- appDto.InsertProduct{
				Name:  product.Name,
				Desc:  product.Desc,
				Price: product.Price,
			}:
			case <-time.After(2500 * time.Millisecond):
				return rabbitmq.NackRequeue
			}

			return rabbitmq.Ack
		})

		if err != nil {
			log.Fatalf("failed to start consumer")
		}
	}()

	// -------------------------------------------------------------------

	instance := RabbitMsgs{
		conn:      conn,
		publisher: publisher,
		consumer:  consumer,

		exchangeName:                   deps.ExchangeName,
		queueNameExternalAddedProducts: deps.QueueNameExternalAddedProducts,
		chanForProductAddedExternally:  externalAddedProductsChan,
	}

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

func (r *RabbitMsgs) GetChanForProductAddedExternally(ctx context.Context) (
	<-chan appDto.InsertProduct, error,
) {
	productsChan := r.chanForProductAddedExternally
	return productsChan, nil
}
