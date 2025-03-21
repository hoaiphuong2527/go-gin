package controllers

import (
	"go-gin-framework/constants"
	dto "go-gin-framework/dto"
	"go-gin-framework/services"
	"go-gin-framework/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Code:    constants.ErrInvalidRequest,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "Fetch users successfully",
		Data:    users,
	})
}

func CreateUser(c *gin.Context) {
	var userDTO dto.CreateUserDTO

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Code:    constants.ErrInvalidRequest,
			Message: err.Error(),
		})
		return
	}

	user, err := services.CreateUser(userDTO)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			c.JSON(http.StatusConflict, dto.Response{
				Success: false,
				Code:    appErr.GetCode(),
				Message: appErr.GetMessage(),
			})
		} else {
			c.JSON(http.StatusNotFound, dto.Response{
				Success: false,
				Code:    constants.ErrInvalidRequest,
				Message: "Cannot create user",
			})
		}
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "User created successfully",
		Data:    user,
	})
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := services.GetOne(id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.Response{
			Success: false,
			Code:    constants.ErrUserNotFound,
			Message: "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "User retrieved successfully",
		Data:    user,
	})
}

func UpdateUser(c *gin.Context) {
	var userDTO dto.UpdateUserDTO
	id := c.Param("id")
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Code:    constants.ErrInvalidRequest,
			Message: err.Error(),
		})
		return
	}

	user, err := services.UpdateUser(id, userDTO)
	if err != nil {
		if err.Error() == "DUPLICATE_USER" {
			c.JSON(http.StatusConflict, dto.Response{
				Success: false,
				Code:    constants.ErrDuplicateUser,
				Message: "User already exists",
			})
		} else {
			c.JSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Code:    constants.ErrDatabaseError,
				Message: "Cannot update user",
			})
		}
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "User updated successfully",
		Data:    user,
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Code:    constants.ErrDatabaseError,
			Message: "Cannot delete user",
		})
		return
	}

	var response dto.UserResponseDTO
	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "User deleted successfully",
		Data:    response,
	})
}
