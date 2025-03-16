package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"runtime/debug"
)

// 패닉 핸들러 (unchecked Exception 같이 서버가 꺼질만한게 발생하면 Stack Trace 출력)
func RecoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("Panic occurred:", r)
				log.Println(string(debug.Stack())) // Stack Trace 출력
				c.JSON(500, gin.H{"error": "Internal Server Error"})
			}
		}()
		c.Next()
	}
}
