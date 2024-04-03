package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func backgroundWork(ctxForCancel context.Context) {
	for {
		select {
		case <-ctxForCancel.Done():
			fmt.Println("cancel")
			return
		case <-time.After(1000 * time.Millisecond):
			fmt.Println(".")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(
		context.Background(),
	)

	go backgroundWork(ctx)
	go backgroundWork(ctx)
	go backgroundWork(ctx)
	//...

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Приложение запущено. Для завершения нажмите Ctrl + C")

	<-sigCh
	cancel()

	time.Sleep(5 * time.Second)
	//...

	fmt.Println("Получен сигнал завершения. Выход из приложения.")
}
