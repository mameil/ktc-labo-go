package router

import (
	"github.com/gin-gonic/gin"
	"kct-labo-go/kct-labo-go/controller"
	"kct-labo-go/kct-labo-go/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())      //이게 뭐야
	r.Use(middleware.Logger()) //이게 뭐야

	api := r.Group("/test") //이게 뭐야

	{
		api.GET("/ping", controller.GetPing)
	}

	return r
}
