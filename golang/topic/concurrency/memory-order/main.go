package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

// publish pattern: writer пишет data+ready, reader ждёт ready и читает data.
// без Mutex - data race; с Mutex - happens-before (Unlock → Lock).

func main() {
	useMutex := flag.Bool("mutex", true, "protect shared state with Mutex")
	flag.Parse()

	var (
		mu    sync.Mutex
		data  int
		ready bool
		wg    sync.WaitGroup
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond) // reader успеет стартовать

		if *useMutex {
			mu.Lock()
			data = 42
			ready = true
			mu.Unlock()
			return
		}
		data = 42
		ready = true
	}()

	go func() {
		defer wg.Done()
		for {
			if *useMutex {
				mu.Lock()
				r, d := ready, data
				mu.Unlock()
				if r {
					fmt.Println("got", d, "(mutex=true)")
					return
				}
				continue
			}
			if ready {
				fmt.Println("got", data, "(mutex=false)")
				return
			}
		}
	}()

	wg.Wait()
}
