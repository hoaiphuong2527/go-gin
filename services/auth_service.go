package services

import (
	"go-gin-framework/constants"
	"go-gin-framework/dto"
	"go-gin-framework/middlewares"
	"go-gin-framework/repositories"
	"go-gin-framework/utils"
)

func Login(dtoLogin dto.LoginDTO) (dto.AuthResponseDTO, error) {
	user, err := repositories.GetUserByEmail(dtoLogin.Email)
	if err != nil {
		return dto.AuthResponseDTO{}, utils.NewAppError(constants.ErrDatabaseError, err.Error())
	}
	if check := user.CheckPassword(dtoLogin.Password); !check {
		return dto.AuthResponseDTO{}, utils.NewAppError(constants.InvalidCredentials, "Invalid credentials")
	}
	token, err := middlewares.GenerateJWT(user.ID)
	if err != nil {
		return dto.AuthResponseDTO{}, utils.NewAppError(constants.ErrGenerateToken, "Cannot generate token")
	}
	var response dto.AuthResponseDTO
	response.Token = token

	return response, nil
}
