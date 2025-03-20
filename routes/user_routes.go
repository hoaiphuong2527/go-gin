package routes

import (
	"go-gin-framework/controllers"
	"go-gin-framework/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", middlewares.AuthMiddleware(), controllers.GetUsers)
		userRoutes.POST("/", controllers.CreateUser)
	}

	return r
}
