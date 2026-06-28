package main

import (
	"fmt"
)

const (
	monday = (iota + 1) * 2
	tuesday
	wednesday
)

func main() {
	fmt.Println(
		monday,
		tuesday,
		wednesday,
	)
}
