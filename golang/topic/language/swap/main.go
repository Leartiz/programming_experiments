package main

import "fmt"

func main() {
	a, b := 1, 2
	fmt.Println("before:", a, b)

	a, b = b, a

	fmt.Println("after:", a, b)
}
