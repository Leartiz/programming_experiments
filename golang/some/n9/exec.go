package n9

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Exec() {
	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	/*
		From Client: GET http://127.0.0.1:8080/user/josh
		---> Hello josh
	*/

	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")

		message := name + " is " + action
		c.String(http.StatusOK, message)
	})
	/*
		From Client: GET http://127.0.0.1:8080/user/josh/send/to/joshua

		---> josh is /send/to/joshua
	*/

	router.Run(":8080")
}
