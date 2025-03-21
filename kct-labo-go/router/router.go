package router

import (
	"github.com/gin-gonic/gin"
	"kct-labo-go/kct-labo-go/controller"
	"kct-labo-go/kct-labo-go/middleware"
	_ "log"
	_ "os"
)

func SetupRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.Logger()) //이게 뭐야
	//r.Use(middleware.KonaLoggingMiddleware())
	r.Use(middleware.RecoverMiddleware()) // 패닉 핸들링 미들웨어 적용

	api := r.Group("/test") //이게 뭐야

	{
		api.GET("/ping", controller.GetPing)
		api.GET("/ping/error", controller.GetPingError)
		api.POST("/pong", controller.PostPong)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, ktc-labo!"})
	})

	return r
}
