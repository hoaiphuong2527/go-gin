package middlewares

import (
	"fmt"
	"go-gin-framework/constants"
	"go-gin-framework/dto"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: No .env file found")
	}
}

var secretKey = []byte(os.Getenv("JWT_SECRET"))

// GenerateJWT tạo token JWT
func GenerateJWT(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(secretKey)
}

// JWTMiddleware authenticate JWT
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get token from Header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, dto.AppResponse{
				Success: false,
				Code:    constants.NotFoundToken,
				Message: "Token not provided",
			})
			c.Abort()
			return
		}

		// trim "Bearer " out token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, dto.AppResponse{
				Success: false,
				Code:    constants.TokenFomartInvalid,
				Message: "Invalid token format",
			})
			c.Abort()
			return
		}

		// decode token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, dto.AppResponse{
				Success: false,
				Code:    constants.InvalidToken,
				Message: "Invalid token",
			})
			c.Abort()
			return
		}

		// save user_id into context to use after handler
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("user_id", uint(claims["user_id"].(float64)))
		} else {
			c.JSON(http.StatusUnauthorized, dto.AppResponse{
				Success: false,
				Code:    constants.InvalidTokenClaims,
				Message: "Invalid token claims",
			})
			c.Abort()
		}

		c.Next()
	}
}
