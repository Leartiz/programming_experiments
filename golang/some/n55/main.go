package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// work выполняет работу, которая занимает 100 мc
func work(ctx context.Context) {
	select {
	case <-time.After(100 * time.Millisecond):
	case <-ctx.Done():
		fmt.Printf("work expired\n")
	}
}

func main() {
	// контекст отменяется по таймауту через 50 мc
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	// после отмены контекста вызывается cleanup
	// будет вызвана в отдельной горутине
	stopCleanup := context.AfterFunc(ctx, cleanup)

	work(ctx)

	stopCleanup() // Если work функция
	_ = 0         // уже начала выполняться -
	_ = 0         // 	отвязать ее не получится?

	//work(ctx)

	cancel()

	// ***

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
}

// cleanup освобождает занятые ресурсы
func cleanup() {
	fmt.Println("cleanup")
}
