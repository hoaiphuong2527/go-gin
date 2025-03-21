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
		c.JSON(http.StatusBadRequest, dto.AppResponse{
			Success: false,
			Code:    constants.ErrInvalidRequest,
			Message: err.Error(),
		})
		return
	}

	token, err := services.Login(loginDto)
	if err != nil {
		utils.HandleErrorAuth(c, err, "Invalid credentials")
		return
	}
	var response dto.AuthResponseDTO
	copier.Copy(&response, &token)

	c.JSON(http.StatusCreated, dto.AppResponse{
		Success: true,
		Message: "",
		Data:    response,
	})
}
