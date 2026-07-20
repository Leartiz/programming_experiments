package main

import (
	"fmt"
	"sync"
)

// done
// -----------------------------------------------------------------------

func merge(cs ...<-chan int) <-chan int {
	resultCh := make(chan int)
	csCount := len(cs)
	if csCount == 0 {
		close(resultCh)
		return resultCh
	}

	done := make(chan struct{})
	go func() {
		for range done {
			csCount -= 1
			if csCount == 0 {
				break
			}
		}

		close(resultCh)
		close(done)
	}()

	for _, c := range cs {
		go func(c <-chan int) {
			for val := range c {
				resultCh <- val
			}

			done <- struct{}{}
		}(c)
	}

	// ограничим интерфейс
	return resultCh
}

// WaitGroup
// -----------------------------------------------------------------------

func mergeWithWg(cs ...<-chan int) <-chan int {
	resultCh := make(chan int)
	if len(cs) == 0 {
		close(resultCh)
		return resultCh
	}

	// ***

	var wg sync.WaitGroup
	wg.Add(len(cs))

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for _, c := range cs {
		go func(c <-chan int) {
			for val := range c {
				resultCh <- val
			}
			wg.Done()
		}(c)
	}

	return resultCh
}

// -----------------------------------------------------------------------

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		defer close(a)
		a <- 1
		a <- 2
	}()

	go func() {
		defer close(b)
		b <- 3
		b <- 5
	}()

	go func() {
		defer close(c)
		c <- 1001
		c <- 10101
	}()

	out := mergeWithWg(a, b, c)
	for v := range out {
		fmt.Println(v)
	}

	// ***

	{
		out := merge()
		for v := range out {
			fmt.Println(v)
		}
	}
}
