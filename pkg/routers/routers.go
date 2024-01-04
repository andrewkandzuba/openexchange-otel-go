package routers

import (
	"openexchange-otel-go/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) *gin.Engine {

    r.GET("/hello", handlers.Greetings())

	m := r.Group("/probs/")
	{
		m.GET("/healthz", handlers.Liveness())
		m.GET("/readyz", handlers.Readiness())
	}

    return r
}