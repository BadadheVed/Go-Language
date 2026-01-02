package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func Limiter(limiter *rate.Limiter) gin.HandlerFunc {

	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit hitted"})
			return
		}

		c.Next()

	}

}
