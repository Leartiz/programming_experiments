package main

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound = errors.New("not found")
)

func main() {
	err := fmt.Errorf("database connection failed: %w", ErrNotFound)
	fmt.Printf("%v\n", err)

}
