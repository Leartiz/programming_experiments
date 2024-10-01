package main

import (
	"fmt"
)

// -----------------------------------------------------------------------

type Avg[T int | float64] struct {
	sum   T
	count int
}

func (a *Avg[T]) Add(value T) *Avg[T] {
	a.sum += value
	a.count += 1
	return a
}

func (a *Avg[T]) Val() T {
	if a.count == 0 {
		return 0
	}

	return a.sum / T(a.count)
}

// -----------------------------------------------------------------------

func main() {
	intAvg := Avg[int]{}
	intAvg.Add(4).Add(3).Add(2)
	fmt.Println(intAvg.Val()) // 3

	floatAvg := Avg[float64]{}
	floatAvg.Add(4.0).Add(3.0)
	floatAvg.Add(2.0).Add(1.0)
	fmt.Println(floatAvg.Val()) // 2.5
}
