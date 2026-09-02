package main

import (
	"fmt"
	"time"

	"golang/interview/ttl-cache/cache"
)

func main() {
	c := cache.NewCache(500 * time.Millisecond)

	c.Set("hello", "world", 2*time.Second)

	if v, ok := c.Get("hello"); ok {
		fmt.Println(v)
	}

	fmt.Println("size:", c.Size())
}
