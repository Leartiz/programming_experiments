package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	pingPongCount := 0
	engine.GET("/ping", func(ctx *gin.Context) {
		pingPongCount += 1
		ctx.String(http.StatusOK, "pong %v", pingPongCount)
	})

	http.ListenAndServe(
		":54000", engine)
}
