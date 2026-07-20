package main

import "fmt"

func main() {
	var ch chan int // nil

	// send / recv на nil-канале блокируются навсегда:
	//   <-ch
	//   ch <- 1

	select {
	case <-ch:
	case ch <- 1:
	default:
		fmt.Println("nil: neither send nor recv ready")
	}
}
