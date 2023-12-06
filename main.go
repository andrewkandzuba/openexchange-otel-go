package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
    "openexchange-otel-go/pkg/utils"
)

const ENV_PORT = "PORT"

func main() {
	r := gin.Default()

	// Apply middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Define routes
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin!")
	})

	// Start the server
	port := utils.GetEnv(ENV_PORT, "8080")
	r.Run(":" + port)
}
