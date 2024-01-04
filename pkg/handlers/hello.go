package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Greetings() gin.HandlerFunc {
    return func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin!")
	}
}
