package utils

import (
	"go-gin-framework/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error, defaultMsg string, codeHttp ...int) {
	if appErr, ok := err.(*AppError); ok {
		code := http.StatusConflict
		if len(codeHttp) > 0 {
			code = codeHttp[0]
		}
		c.JSON(code, dto.AppResponse{
			Success: false,
			Message: appErr.Error(),
		})
	} else {
		c.JSON(http.StatusInternalServerError, dto.AppResponse{
			Success: false,
			Message: defaultMsg,
		})
	}
}

func HandleErrorAuth(c *gin.Context, err error, defaultMsg string) {
	if appErr, ok := err.(*AppError); ok {
		c.JSON(http.StatusForbidden, dto.AppResponse{
			Success: false,
			Message: appErr.Error(),
		})
	} else {
		c.JSON(http.StatusForbidden, dto.AppResponse{
			Success: false,
			Message: defaultMsg,
		})
	}
}
