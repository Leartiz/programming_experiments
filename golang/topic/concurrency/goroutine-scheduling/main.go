package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		var u int = 0
		for {
			u -= 2
			if u == 1 {
				break
			}
		}
	}()

	fmt.Printf("before")
	<-time.After(time.Second)

	// "никогда" не выполнится, хотя версия go 1.21.6...
	fmt.Println("go 1.13 has never been here")
}
