package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func run(ctx context.Context) {
	wait := make(chan struct{}, 1)

	go func() {
		defer func() {
			wait <- struct{}{}
		}()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Break the loop")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("Hello in a loop")
			}
		}
	}()

	go func() {
		defer func() {
			wait <- struct{}{}
		}()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Break the loop")
				return
			case <-time.After(1 * time.Second):
				fmt.Println("Ciao in a loop")
			}
		}
	}()

	// wait for two goroutines to finish
	<-wait
	<-wait

	fmt.Println("Main done")
}

func main() {
	fmt.Println("with_chan")

	rootCtx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt, syscall.SIGTERM)
	defer stop()

	run(rootCtx)
}
