package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() func(c *gin.Context) {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		log.Printf("[GIN] %s - %s %s (%v)", c.ClientIP(), c.Request.Method, c.Request.URL, duration)
	}
}
