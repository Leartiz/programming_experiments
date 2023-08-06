package n1

import (
	"context"
	"fmt"
	"log"
	"time"
)

type key int

const (
	userIdCtx key = 0
)

func Exec() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, userIdCtx, 1)
	ctx, cancel := context.WithCancel(ctx)

	// ***

	go func() {
		fmt.Scanln()
		cancel()
	}()

	// ***

	processLongTask(ctx)
}

func processLongTask(ctx context.Context) {
	id := ctx.Value(userIdCtx)

	select {
	case <-time.After(2 * time.Second):
		fmt.Println("done processing", id)
	case <-ctx.Done():
		log.Println("request canceled")
	}
}
