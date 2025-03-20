package routes

import (
	"go-gin-framework/controllers"
	// "go-gin-framework/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	userGroup := r.Group("/users")
	{
		// userGroup.GET("/", middlewares.AuthMiddleware(), controllers.GetAllUsers)

		userGroup.POST("/", controllers.CreateUser)
		userGroup.GET("/", controllers.GetAllUsers)
		userGroup.GET("/:id", controllers.GetUser)
		userGroup.PUT("/:id", controllers.UpdateUser)
		userGroup.DELETE("/:id", controllers.DeleteUser)
	}

	return r
}
