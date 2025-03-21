package middlewares

import (
	"go-gin-framework/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists || userRole != requiredRole {
			c.JSON(http.StatusForbidden, dto.AppResponse{
				Success: false,
				Message: "Access denied",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
