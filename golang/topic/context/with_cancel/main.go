package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	rootCtx, rootStop := signal.NotifyContext(context.Background(),
		os.Interrupt, // SIGINT
		syscall.SIGTERM,
	)
	defer rootStop()

	var ctx context.Context
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(rootCtx, 5*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		counter := 0

		for {
			select {
			case <-ctx.Done():
				fmt.Println("routine stopped!")
				return
			case <-time.After(1 * time.Second):
				fmt.Println(counter)
				counter++
			}
		}
	}(ctx)

	<-ctx.Done()

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("run duration timeout reached, shutting down scheduler")
	} else {
		fmt.Println("received stop signal, shutting down scheduler")
	}
}
