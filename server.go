package main

import (
	"net/http"

	"github.com/dwilkolek/go-proto/def"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Colors
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/pingproto", func(c *gin.Context) {
		data := &def.Message{
			Body: "Pong",
			Ttl:  69,
		}

		c.ProtoBuf(http.StatusOK, data)
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
