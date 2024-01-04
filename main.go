package main

import (
	"openexchange-otel-go/pkg/routers"
	"openexchange-otel-go/pkg/utils"

	"github.com/gin-gonic/gin"
)

const ENV_PORT = "PORT"

func main() {
	r := gin.Default()

	// Apply middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Define routes
	r = routers.SetupRouter(r)


	// Start the server
	port := utils.GetEnv(utils.ENV_PORT, "8080")
	r.Run(":" + port)
}
