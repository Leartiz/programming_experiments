package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "dfly:6379",
		Password: "",
		DB:       0,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// ***

	engine := gin.Default()
	engine.GET("/ping", func(ctx *gin.Context) {
		err := rdb.Incr(ctx, "ping").Err()
		if err != nil {
			fmt.Printf("%v\n", err)
			ctx.Status(http.StatusInternalServerError)
		}

		val, err := rdb.Get(ctx, "ping").Result()
		if err != nil {
			fmt.Printf("%v\n", err)
			ctx.Status(http.StatusInternalServerError)
		}

		ctx.String(http.StatusOK, "pong %v", val)
	})

	http.ListenAndServe(
		":54545", engine)
}
