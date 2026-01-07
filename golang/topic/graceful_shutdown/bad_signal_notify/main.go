package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	// This will not work as expected!
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		stopped := false
		for !stopped {
			select {
			case <-exit: // Only one go routine will get the termination signal
				fmt.Println("Break the loop: hello")
				stopped = true
			case <-time.After(1 * time.Second):
				fmt.Println("Hello in a loop")
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		stopped := false
		for !stopped {
			select {
			case <-exit: // Only one go routine will get the termination signal
				fmt.Println("Break the loop: ciao")
				stopped = true
			case <-time.After(1 * time.Second):
				fmt.Println("Ciao in a loop")
			}
		}
	}()

	wg.Wait()
	fmt.Println("Main done")
}
