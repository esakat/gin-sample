package router

import (
	"github.com/esakat/gin-sample/app"
	"github.com/gin-gonic/gin"
	"github.com/esakat/gin-sample/service"
)

func Init(app *app.App) *gin.Engine {

	router := gin.Default()

	router.Use(app.AppendEnv)

	v1 := router.Group("/v1")
	{
		v1.GET("/user", service.GetUsers)
		v1.GET("/user/:id", service.GetUser)
		v1.POST("/user", service.PostUser)
	}

	return router
}
