package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch)

	// из закрытого канала читается zero value, не блокируется
	fmt.Println(<-ch, <-ch) // 0 0

	v, ok := <-ch
	fmt.Println(v, ok) // 0 false
}
