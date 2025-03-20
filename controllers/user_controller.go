package controllers

import (
	"go-gin-framework/constants"
	dto "go-gin-framework/dto"
	"go-gin-framework/services"
	"go-gin-framework/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
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

	var response dto.UserResponseDTO
	copier.Copy(&response, &users)

	c.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "Fetch users successfully",
		Data:    response,
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

	var response dto.UserResponseDTO
	copier.Copy(&response, &user)

	c.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "User created successfully",
		Data:    response,
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

	var response dto.UserResponseDTO
	copier.Copy(&response, &user)

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "User retrieved successfully",
		Data:    response,
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
	var response dto.UserResponseDTO
	copier.Copy(&response, &user)

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "User updated successfully",
		Data:    response,
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
