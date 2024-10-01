package main

import (
	"fmt"
)

// Pair - пара значений.
type Pair[T any] struct {
	first  T
	second T
}

type Number[T int | float64] struct {
	val T
}

// Swap меняет местами значения в паре.
func (p *Pair[T]) Swap() {
	p.first, p.second = p.second, p.first
}

func main() {
	// срез целых чисел
	intSlice := []int{1, 2, 3}

	// срез строк
	strSlice := []string{"a", "b", "c"}

	fmt.Printf("%v\n", intSlice)
	fmt.Printf("%v\n", strSlice)

	// ***

	intPair := Pair[int]{5, 3}
	fmt.Println(intPair)
	intPair.Swap()
	fmt.Println(intPair)
	// {3 5}

	strPair := Pair[string]{"one", "two"}
	fmt.Println(strPair)
	strPair.Swap()
	fmt.Println(strPair)
	// {two one}
}
