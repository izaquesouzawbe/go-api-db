package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func authorize(expectedToken string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader != expectedToken {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
