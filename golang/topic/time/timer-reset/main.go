package main

import (
	"fmt"
	"time"

	"math/rand"
)

func main() {
	nums := make(chan int)
	done := make(chan struct{})

	// печатает числа из канала nums.
	go func() {
		defer close(done)
		timeout := 10 * time.Millisecond
		timer := time.NewTimer(timeout)

		for {
			select {
			case n, ok := <-nums:
				if !ok {
					return
				}
				fmt.Print(n)

			case <-timer.C:
				fmt.Print("!")
			}

			reset(timer, timeout)
		}
	}()

	// отправляет числа в канал nums
	// с задержкой от 5 до 15 миллисекунд.
	for i := range 10 {
		delay := time.Duration(5+rand.Intn(11)) * time.Millisecond
		time.Sleep(delay)
		nums <- i
	}

	close(nums)
	<-done
}

// reset сбрасывает таймер и перезапускает его.
func reset(timer *time.Timer, timeout time.Duration) {
	/*
		Stop prevents the Timer from firing.
		It returns true if the call stops the timer,
			false if the timer has already expired or been stopped.
	*/
	if !timer.Stop() {
		// таймер уже сработал,
		// очищаем его канал
		select {
		case <-timer.C:
		default:
		}
	}
	timer.Reset(timeout)
}
