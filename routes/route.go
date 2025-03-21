package routes

import (
	"go-gin-framework/controllers"
	"go-gin-framework/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	apiGroup := r.Group("/api")
	authGroup := apiGroup.Group("/auth")
	{
		authGroup.POST("/login", controllers.Login)
	}

	meGroup := apiGroup.Group("/me")
	meGroup.Use(middlewares.JWTMiddleware())
	meGroup.GET("/", controllers.GetUserProfile)

	userGroup := apiGroup.Group("/users")
	userGroup.Use(middlewares.JWTMiddleware())
	userGroup.POST("/", middlewares.RequireRole("admin"), controllers.CreateUser)
	userGroup.GET("/", middlewares.RequireRole("admin"), controllers.GetAllUsers)
	userGroup.GET("/:id", controllers.GetUser)
	userGroup.PUT("/:id", controllers.UpdateUser)
	userGroup.DELETE("/:id", controllers.DeleteUser)

	return r
}
