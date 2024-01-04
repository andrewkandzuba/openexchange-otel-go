package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Liveness() gin.HandlerFunc {
    return func(c *gin.Context) {
		c.String(http.StatusOK, "Healthy!")
	}
}

func Readiness() gin.HandlerFunc {
    return func(c *gin.Context) {
		c.String(http.StatusOK, "Ready!")
	}
}

