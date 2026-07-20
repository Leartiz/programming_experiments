package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== 1) baseline ===")
	fmt.Printf("NumCPU=%d GOMAXPROCS=%d NumGoroutine=%d\n",
		runtime.NumCPU(), runtime.GOMAXPROCS(0), runtime.NumGoroutine())

	fmt.Println("\n=== 2) many goroutines (mostly sleeping) ===")
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			time.Sleep(20 * time.Millisecond)
			fmt.Printf("  sleeper %d done\n", id)
		}(i)
	}
	wg.Wait()

	fmt.Println("\n=== 3) GOMAXPROCS=1, 3 goroutines interleave on Sleep ===")
	runtime.GOMAXPROCS(1)
	done := make(chan struct{})
	for i := 0; i < 3; i++ {
		go func(id int) {
			for step := 0; step < 3; step++ {
				fmt.Printf("  g%d step %d\n", id, step)
				time.Sleep(5 * time.Millisecond) // yield point
			}
			done <- struct{}{}
		}(i)
	}
	for i := 0; i < 3; i++ {
		<-done
	}

	fmt.Println("\n=== 4) unbuffered chan — rendezvous (G blocks until pair) ===")
	ch := make(chan int)
	go func() {
		fmt.Println("  sender: before send")
		ch <- 42
		fmt.Println("  sender: after send")
	}()
	time.Sleep(10 * time.Millisecond)
	fmt.Println("  main: before receive")
	fmt.Println("  main: got", <-ch)

	fmt.Println("\n=== 5) cpu spin + sleeper (preemption since Go 1.14) ===")
	finished := make(chan struct{})
	go func() {
		// CPU-bound; runtime eventually preempts (~10ms)
		spin := 0
		for spin < 50_000_000 {
			spin++
		}
		fmt.Println("  spinner finished")
		close(finished)
	}()
	go func() {
		time.Sleep(1 * time.Millisecond)
		fmt.Println("  sleeper ran while spinner was spinning")
	}()
	<-finished

	runtime.GOMAXPROCS(runtime.NumCPU()) // restore
	fmt.Println("\n=== done ===")
}
