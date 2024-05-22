package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	uri          = "amqp://guest:guest@localhost:5672/"
	exchangeName = "boiler_requests"
	queueName    = "service"
	routingKey   = "hhhhh"
)

var (
	body       = flag.String("body", "hi", "Body of message")
	continuous = flag.Bool("continuous", true, "Keep publishing messages at a 1msg/sec rate")

	InfoLog = log.New(os.Stdout, "[INFO] ", log.LstdFlags|log.Lmsgprefix)
	WarnLog = log.New(os.Stderr, "[WARNING] ", log.LstdFlags|log.Lmsgprefix)
	ErrLog  = log.New(os.Stderr, "[ERROR] ", log.LstdFlags|log.Lmsgprefix)
)

func init() {
	flag.Parse()
}

func main() {
	exitCh := make(chan struct{})

	confirmsCh := make(chan *amqp.DeferredConfirmation)
	confirmsDoneCh := make(chan struct{})
	publishOkCh := make(chan struct{}, 1)

	setupCloseHandler(exitCh)

	startConfirmHandler(publishOkCh, confirmsCh, confirmsDoneCh, exitCh)
	publish(publishOkCh, confirmsCh, confirmsDoneCh, exitCh)
}

func setupCloseHandler(exitCh chan struct{}) {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		InfoLog.Printf("close handler: Ctrl+C pressed in Terminal")
		close(exitCh)
	}()
}

func publish(publishOkCh <-chan struct{},
	confirmsCh chan<- *amqp.DeferredConfirmation,
	confirmsDoneCh <-chan struct{},
	exitCh chan struct{},
) {

	config := amqp.Config{
		Vhost:      "/",
		Properties: amqp.NewConnectionProperties(),
	}
	config.Properties.SetClientConnectionName("producer-with-confirms")

	// create connection

	conn, err := amqp.DialConfig(uri, config)
	if err != nil {
		ErrLog.Fatalf("producer: error in dial: %s", err)
	}
	defer conn.Close()

	// create channel

	channel, err := conn.Channel()
	if err != nil {
		ErrLog.Fatalf("error getting a channel: %s", err)
	}
	defer channel.Close()

	if err := channel.QueueBind(queueName, routingKey, exchangeName, false, nil); err != nil {
		ErrLog.Fatalf("producer: Queue Bind: %s", err)
	}

	InfoLog.Printf("producer: enabling publisher confirms.")
	if err := channel.Confirm(false); err != nil {
		ErrLog.Fatalf("producer: channel could not be put into confirm mode: %s", err)
	}

	for {
		canPublish := false
		InfoLog.Println("producer: waiting on the OK to publish...")
		for {
			select {
			case <-confirmsDoneCh:
				InfoLog.Println("producer: stopping, all confirms seen")
				return
			case <-publishOkCh:
				InfoLog.Println("producer: got the OK to publish")
				canPublish = true
			case <-time.After(time.Second):
				WarnLog.Println("producer: still waiting on the OK to publish...")
				continue
			}
			if canPublish {
				break
			}
		}

		InfoLog.Printf("producer: publishing %dB body (%q)", len(*body), *body)
		dConfirmation, err := channel.PublishWithDeferredConfirm(
			exchangeName,
			routingKey,
			true,
			false,
			amqp.Publishing{
				Headers:         amqp.Table{},
				ContentType:     "text/plain",
				ContentEncoding: "",
				DeliveryMode:    amqp.Persistent,
				Priority:        0,
				AppId:           "sequential-producer",
				Body:            []byte(*body),
			},
		)
		if err != nil {
			ErrLog.Fatalf("producer: error in publish: %s", err)
		}

		select {
		case <-confirmsDoneCh:
			InfoLog.Println("producer: stopping, all confirms seen")
			return
		case confirmsCh <- dConfirmation:
			InfoLog.Println("producer: delivered deferred confirm to handler")
		}

		select {
		case <-confirmsDoneCh:
			InfoLog.Println("producer: stopping, all confirms seen")
			return
		case <-time.After(time.Second):
			if *continuous {
				continue
			} else {
				InfoLog.Println("producer: initiating stop")
				close(exitCh)
				select {
				case <-confirmsDoneCh:
					InfoLog.Println("producer: stopping, all confirms seen")
					return
				case <-time.After(time.Second * 10):
					WarnLog.Println("producer: may be stopping with outstanding confirmations")
					return
				}
			}
		}
	}
}

func startConfirmHandler(publishOkCh chan<- struct{},
	confirmsCh <-chan *amqp.DeferredConfirmation,
	confirmsDoneCh chan struct{},
	exitCh <-chan struct{},
) {
	go func() {
		confirms := make(map[uint64]*amqp.DeferredConfirmation)

		for {
			select {
			case <-exitCh:
				exitConfirmHandler(confirms, confirmsDoneCh)
				return
			default:
			}

			// ***

			outstandingConfirmationCount := len(confirms)

			// Note: 8 is arbitrary, you may wish to allow more outstanding confirms before blocking publish
			if outstandingConfirmationCount <= 8 {
				select {
				case publishOkCh <- struct{}{}:
					InfoLog.Println("confirm handler: sent OK to publish")
				case <-time.After(time.Second * 5):
					WarnLog.Println("confirm handler: timeout indicating OK to publish (this should never happen!)")
				}
			} else {
				WarnLog.Printf("confirm handler: waiting on %d outstanding confirmations, blocking publish", outstandingConfirmationCount)
			}

			select {
			case confirmation := <-confirmsCh:
				deliveryTag := confirmation.DeliveryTag
				confirms[deliveryTag] = confirmation
			case <-exitCh:
				exitConfirmHandler(confirms, confirmsDoneCh)
				return
			}

			checkConfirmations(confirms)
		}
	}()
}

func exitConfirmHandler(confirms map[uint64]*amqp.DeferredConfirmation, confirmsDoneCh chan struct{}) {
	InfoLog.Println("confirm handler: exit requested")
	waitConfirmations(confirms)
	close(confirmsDoneCh)
	InfoLog.Println("confirm handler: exiting")
}

func checkConfirmations(confirms map[uint64]*amqp.DeferredConfirmation) {
	InfoLog.Printf("confirm handler: checking %d outstanding confirmations", len(confirms))
	for k, v := range confirms {
		if v.Acked() {
			InfoLog.Printf("confirm handler: confirmed delivery with tag: %d", k)
			delete(confirms, k)
		}
	}
}

func waitConfirmations(confirms map[uint64]*amqp.DeferredConfirmation) {
	InfoLog.Printf("confirm handler: waiting on %d outstanding confirmations", len(confirms))

	checkConfirmations(confirms)

	for k, v := range confirms {
		select {
		case <-v.Done():
			InfoLog.Printf("confirm handler: confirmed delivery with tag: %d", k)
			delete(confirms, k)
		case <-time.After(time.Second):
			WarnLog.Printf("confirm handler: did not receive confirmation for tag %d", k)
		}
	}

	outstandingConfirmationCount := len(confirms)
	if outstandingConfirmationCount > 0 {
		ErrLog.Printf("confirm handler: exiting with %d outstanding confirmations", outstandingConfirmationCount)
	} else {
		InfoLog.Println("confirm handler: done waiting on outstanding confirmations")
	}
}
