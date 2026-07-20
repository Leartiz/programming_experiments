package main

import (
	"fmt"
	"time"
)

// буфер откладывает точку синхронизации:
// unbuffered - зависание сразу на первом send
// buffered(1) - первый send ок, deadlock на втором

func trySends(label string, ch chan int, n int) {
	done := make(chan struct{})
	go func() {
		for i := 0; i < n; i++ {
			ch <- i
			fmt.Println(label, "sent", i)
		}
		close(done)
	}()

	select {
	case <-done:
		fmt.Println(label, "ok")
	case <-time.After(100 * time.Millisecond):
		fmt.Println(label, "stuck (would deadlock)")
	}
}

func main() {
	trySends("unbuffered", make(chan int), 1)
	trySends("buffered", make(chan int, 1), 2)
}
