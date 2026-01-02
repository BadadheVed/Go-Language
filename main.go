package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func Limiter(limiter *rate.Limiter) gin.HandlerFunc {

	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(
				http.StatusTooManyRequests,
				gin.H{"error": "rate limit hit"},
			)
			return
		}

		c.Next()

	}

}
func main() {
	r := gin.Default()

	limit := rate.NewLimiter(rate.Limit(10.0/60.0), 10)

	r.Use(Limiter(limit))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":3000")
}
