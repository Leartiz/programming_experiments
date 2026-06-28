package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type key int

const userIDCtx key = 0

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, userIDCtx, 1)
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		fmt.Scanln()
		cancel()
	}()

	processLongTask(ctx)
}

func processLongTask(ctx context.Context) {
	id := ctx.Value(userIDCtx)

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("done processing", id)
	case <-ctx.Done():
		log.Println("request canceled")
	}
}
