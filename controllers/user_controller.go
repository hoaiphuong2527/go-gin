package controllers

import (
	dto "go-gin-framework/dto"
	"go-gin-framework/services"
	"go-gin-framework/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUsers(c)
	if err != nil {
		utils.HandleError(c, err, "Failed to fetch users")
		return
	}

	c.JSON(http.StatusOK, dto.AppResponse{
		Success: true,
		Message: "Fetched users successfully",
		Data:    users,
	})
}

func CreateUser(c *gin.Context) {
	var userDTO dto.CreateUserDTO

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		utils.HandleError(c, err, "Invalid request data")
		return
	}

	user, err := services.CreateUser(userDTO)
	if err != nil {
		utils.HandleError(c, err, "Cannot create user")
		return
	}

	c.JSON(http.StatusCreated, dto.AppResponse{
		Success: true,
		Message: "User created successfully",
		Data:    user,
	})
}

func GetUser(c *gin.Context) {
	user, err := services.GetOne(c.Param("id"))
	if err != nil {
		utils.HandleError(c, err, "User not found")
		return
	}

	c.JSON(http.StatusOK, dto.AppResponse{
		Success: true,
		Message: "User retrieved successfully",
		Data:    user,
	})
}

func UpdateUser(c *gin.Context) {
	var userDTO dto.UpdateUserDTO
	id := c.Param("id")

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		utils.HandleError(c, err, "Invalid request data")
		return
	}

	user, err := services.UpdateUser(id, userDTO)
	if err != nil {
		utils.HandleError(c, err, "Cannot update user")
		return
	}

	c.JSON(http.StatusOK, dto.AppResponse{
		Success: true,
		Message: "User updated successfully",
		Data:    user,
	})
}

func DeleteUser(c *gin.Context) {
	if err := services.DeleteUser(c.Param("id")); err != nil {
		utils.HandleError(c, err, "Cannot delete user")
		return
	}

	c.JSON(http.StatusOK, dto.AppResponse{
		Success: true,
		Message: "User deleted successfully",
	})
}
