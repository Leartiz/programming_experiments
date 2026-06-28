package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func printAboutRedisClient(client *redis.Client) {
	stats := client.PoolStats()
	fmt.Println("Hits:", stats.Hits)
	fmt.Println("Misses:", stats.Misses)
	fmt.Println("Stale Conns:", stats.StaleConns)
	fmt.Println("Idle Conns:", stats.IdleConns)
	fmt.Println("Timeouts:", stats.Timeouts)
	//...
}

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	printAboutRedisClient(rdb)

	// ***

	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	// Output: key value
	// key2 does not exist
}
