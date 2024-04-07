package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	ctxWithCancel, cancel := context.WithCancel(context.Background())
	defer cancel()

	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-ctxWithCancel.Done():
			w.Write([]byte("server dropped by cancel"))

		case <-r.Context().Done():
			w.Write([]byte("server dropped")) // does not work!

		case <-time.After(5000 * time.Millisecond):
			w.Write([]byte("pong"))
		}
	})

	svr := &http.Server{
		Addr:           ":57002",
		MaxHeaderBytes: 4096,
		ReadTimeout:    25 * time.Second,
		WriteTimeout:   25 * time.Second,
		Handler:        serveMux,
	}

	go func() {
		err := svr.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// ***

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGABRT,
	)

	fmt.Println("Приложение запущено. Для завершения нажмите Ctrl + C")

	<-sigCh
	cancel()
	time.Sleep(1 * time.Second)

	svr.Close()

	fmt.Println("Получен сигнал завершения. Выход из приложения.")
}
