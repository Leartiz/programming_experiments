package main

import (
	"fmt"
	"sync"
	"time"
)

// buffered chan как семафор: cap = лимит параллелизма

func main() {
	const limit = 2
	sem := make(chan struct{}, limit)

	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		sem <- struct{}{} // занять слот

		go func(id int) {
			defer wg.Done()
			defer func() { <-sem }() // освободить

			fmt.Println("start", id, time.Now().Format("15:04:05.000"))
			time.Sleep(100 * time.Millisecond)
			fmt.Println("done ", id, time.Now().Format("15:04:05.000"))
		}(i)
	}
	wg.Wait()
}
