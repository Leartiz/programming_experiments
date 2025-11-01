package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func run(ctx context.Context) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
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

	wg.Add(1)
	go func() {
		defer wg.Done()
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

	wg.Wait()

	fmt.Println("Main done")
}

func main() {
	fmt.Println("with_wait_group")

	rootCtx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt, syscall.SIGTERM)
	defer stop()

	run(rootCtx)
}
