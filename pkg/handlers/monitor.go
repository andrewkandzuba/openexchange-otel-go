package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Liveness() gin.HandlerFunc {
    return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Healthy!"})
	}
}

func Readiness() gin.HandlerFunc {
    return func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"message": "Ready!"})
	}
}

