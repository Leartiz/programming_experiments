package main

import (
	"cmp"
	"fmt"
)

func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	{
		a, b := 5, 3
		max := Max(a, b)
		fmt.Println(max) // 5
	}
	{
		a, b := 5.5, 3.3
		max := Max(a, b)
		fmt.Println(max) // 5.5
	}
}
