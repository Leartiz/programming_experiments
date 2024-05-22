package main

import (
	"fmt"
	"n28/internal/launch"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := launch.Launch(); err != nil {
		fmt.Printf("%v\n", err)
	}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh,
		syscall.SIGINT,
		syscall.SIGHUP,
		syscall.SIGTERM,
		syscall.SIGABRT,
	)

	<-sigCh
}
