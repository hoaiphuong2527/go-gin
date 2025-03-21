package controllers

import (
	"go-gin-framework/constants"
	"go-gin-framework/dto"
	"go-gin-framework/services"
	"go-gin-framework/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

func Login(c *gin.Context) {
	var loginDto dto.LoginDTO

	if err := c.ShouldBindJSON(&loginDto); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Code:    constants.ErrInvalidRequest,
			Message: err.Error(),
		})
		return
	}

	token, err := services.Login(loginDto)
	if err != nil {
		if appErr, ok := err.(*utils.AppError); ok {
			c.JSON(http.StatusConflict, dto.Response{
				Success: false,
				Code:    appErr.GetCode(),
				Message: appErr.GetMessage(),
			})
		} else {
			c.JSON(http.StatusForbidden, dto.Response{
				Success: false,
				Code:    constants.InvalidCredentials,
				Message: "Invalid credentials",
			})
		}
		return
	}
	var response dto.AuthResponseDTO
	copier.Copy(&response, &token)

	c.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "",
		Data:    response,
	})
}
