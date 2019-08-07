package routers

import (
	"gin-test-go/configs"
	"gin-test-go/controllers"
	"gin-test-go/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter(c configs.Conf) *gin.Engine {
	router := gin.Default()

	if c.Deploy != "dev" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router.Static("/apidoc", "./resources/apidoc")

	v1 := router.Group("/v1")
	v1.Use(middlewares.Auth())
	{
		v1.GET("/users", controllers.GetUsers)
		v1.POST("/users",controllers.CreateUser)
	}
	return router
}
