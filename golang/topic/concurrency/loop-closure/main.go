package main

import (
	"fmt"
	"time"
)

func main() {
	s := []int{1, 2, 3}

	i := 0
	for i = 0; i < len(s); i++ {
		go func(i int) {
			time.Sleep(1 * time.Millisecond)
			fmt.Print(s[i])
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println()
}
