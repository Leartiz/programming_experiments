package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic:", r)
		}
	}()

	ch <- 1 // send on closed channel
}
