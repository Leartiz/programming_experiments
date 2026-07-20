package main

import (
	"context"
	"flag"
	"fmt"
	"sync"
	"time"
)

const cleanupDelay = 200 * time.Millisecond

func watch(ctx context.Context, onDone func()) {
	<-ctx.Done()
	time.Sleep(cleanupDelay)    // simulate close DB / flush logs
	fmt.Println("cleanup done") // не вызывается!
	if onDone != nil {
		onDone()
	}
}

// -----------------------------------------------------------------------

// BAD: defer cancel on main exit, no wait - process may exit before cleanup
func runBad() {
	fmt.Println("--- bad ---")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go watch(ctx, nil)
	fmt.Println("main returns")
}

// GOOD: cancel + wait until worker finishes cleanup
func runGood() {
	fmt.Println("--- good ---")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan struct{})
	go watch(ctx, func() { close(done) })

	time.Sleep(50 * time.Millisecond) // app runs a bit

	cancel() // no-op
	<-done

	fmt.Println("main returns")
}

// GOOD (variant): sync.WaitGroup instead of done channel
func runGoodWaitGroup() {
	fmt.Println("--- good (WaitGroup) ---")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		watch(ctx, nil)
	}()

	time.Sleep(50 * time.Millisecond)
	cancel()
	wg.Wait()
	fmt.Println("main returns")
}

func main() {
	mode := flag.String("mode", "all", "bad | good | wg | all")
	flag.Parse()

	switch *mode {
	case "bad":
		// обычно нет print (процесс выходит раньше)
		runBad()

	case "good":
		runGood()

	case "wg":
		runGoodWaitGroup()

	case "all":
		runBad()
		fmt.Println()
		runGood()

	default:
		fmt.Println("unknown mode, use: bad | good | wg | all")
	}
}
