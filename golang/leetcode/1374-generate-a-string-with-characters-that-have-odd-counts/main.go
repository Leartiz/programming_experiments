package main

import (
	"fmt"
)

func generateTheString(n int) string {
	bytes := make([]byte, 0, n)
	for i := 0; i < n; i += 1 {
		bytes = append(bytes, 'a')
	}
	if n%2 == 0 {
		bytes[0] = 'b'
	}
	return string(bytes)
}

// v2 через массив частот?

func main() {
	/*
		Input: n = 4
		Output: "pppz" (or any valid string)

		Input: n = 5
		Output: "aaaaa" (or any valid string)

		Input: n = 7
		Output: "jjjjjjj" (or any valid string)
	*/
	{
		fmt.Println(generateTheString(4))
		fmt.Println(generateTheString(5))
		fmt.Println(generateTheString(7))
	}
}
