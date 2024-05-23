package main

import (
	"context"
	"log"
	"n28/internal/launch"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctxForCancel, cancelBkg := context.WithCancel(context.Background())
	ctxForInit, cancelInit := context.WithTimeout(context.Background(), 5*time.Second)

	if err := launch.Launch(ctxForInit, ctxForCancel); err != nil {
		log.Fatalf("%v", err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGABRT,
	)

	<-sigCh
	cancelInit()
	cancelBkg()

	time.Sleep(
		time.Second)
}
