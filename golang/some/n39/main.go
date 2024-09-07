package main

import (
	"fmt"
)

// -----------------------------------------------------------------------

func Reverse[T any](s []T) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

type Person struct {
	Name string
	Age  int
}

// -----------------------------------------------------------------------

func Produce[T any](val T, n int) []T {
	vals := make([]T, n)
	for i := range n {
		vals[i] = val
	}
	return vals
}

func main() {
	{
		intSlice := []int{1, 2, 3}
		fmt.Println(intSlice)
		Reverse(intSlice)
		fmt.Println(intSlice)
		// [3 2 1]
	}
	{
		strSlice := []string{"a", "b", "c"}
		fmt.Println(strSlice)
		Reverse(strSlice) // !
		fmt.Println(strSlice)
		// [c b a]
	}
	{
		personSlice := []Person{
			{Name: "Alice", Age: 10},
			{Name: "Bob", Age: 100},
			{Name: "Cindy", Age: 18},
		}

		Reverse(personSlice)
		fmt.Println(personSlice)
	}
	{
		intSlice := Produce(5, 3)
		fmt.Println(intSlice)

		strSlice := Produce("o", 5)
		fmt.Println(strSlice)
	}
}
