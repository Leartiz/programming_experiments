package main

import (
	"fmt"
	"time"

	"math/rand"
)

func main() {
	var start = time.Now()

	work := func() {
		fmt.Printf("work done after %dms\n",
			time.Since(start).Milliseconds())
	}

	// выполняет функцию work через 10 миллисекунд
	timeout := 10 * time.Millisecond
	start = time.Now()
	t := time.AfterFunc(timeout, work)

	// ждем от 5 до 15 миллисекунд
	delay := time.Duration(5+rand.Intn(11)) * time.Millisecond
	time.Sleep(delay)
	fmt.Printf("%dms has passed...\n", delay.Milliseconds())

	// поведение Reset зависит от того, сработал ли таймер!!!
	t.Reset(timeout)
	start = time.Now()

	time.Sleep(
		50 * time.Millisecond)
}
