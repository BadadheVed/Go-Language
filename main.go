package main

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

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
