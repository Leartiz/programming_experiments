package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Time{}
	fmt.Printf("%v\n", t)
	fmt.Printf("%v\n", t.IsZero())

	t = time.Now()
	fmt.Printf("%v\n", t)
	fmt.Printf("%v\n", t.IsZero())
}
