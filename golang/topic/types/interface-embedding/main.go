package main

import (
	"fmt"
)

type counter interface {
	increment()
	value() uint
}

type simpleCounter struct {
	count uint
}

func (c *simpleCounter) increment() {
	c.count++
}

func (c *simpleCounter) value() uint {
	return c.count
}

type usage struct {
	service string
	counter
}

// -----------------------------------------------------------------------

func main() {
	sc := &simpleCounter{count: 0}

	u := usage{
		service: "api",
		counter: sc,
	}

	u.increment()
	u.increment()
	u.increment()

	fmt.Printf("Service: %s, Count: %d\n", u.service, u.value())

	printCounter(&u)

	// NOTE: можно заменить реализацию counter на другую
	anotherCounter := &simpleCounter{count: 10}
	u2 := usage{
		service: "database",
		counter: anotherCounter,
	}

	fmt.Printf("Service: %s, Count: %d\n",
		u2.service, u2.value())
}

func printCounter(c counter) {
	fmt.Printf("Counter value: %d\n", c.value())
}
