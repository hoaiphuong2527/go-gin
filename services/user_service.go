package services

import (
	"go-gin-framework/constants"
	"go-gin-framework/dto"
	"go-gin-framework/models"
	"go-gin-framework/repositories"
	"go-gin-framework/utils"

	"github.com/jinzhu/copier"
)

func CreateUser(userDTO dto.CreateUserDTO) (dto.UserResponseDTO, error) {
	user := models.User{}
	err := copier.Copy(&user, &userDTO)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}
	if err := user.HashPassword(); err != nil {
		return dto.UserResponseDTO{}, utils.NewAppError(constants.ErrHashPassword, "Failed to hash password")
	}
	if err := repositories.CreateUser(user); err != nil {
		return dto.UserResponseDTO{}, utils.NewAppError(constants.ErrDatabaseError, err.Error())
	}

	var response dto.UserResponseDTO
	copier.Copy(&response, &user)

	return response, nil
}

func GetAllUsers() ([]dto.UserResponseDTO, error) {
	users, errGetList := repositories.GetAllUsers()
	if errGetList != nil {
		return nil, utils.NewAppError(constants.ErrDatabaseError, errGetList.Error())
	}

	var response []dto.UserResponseDTO
	err := copier.Copy(&response, &users)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func GetOne(userId string) (dto.UserResponseDTO, error) {
	user, errDB := repositories.GetUserByID(userId)
	if errDB != nil {
		return dto.UserResponseDTO{}, utils.NewAppError(constants.ErrDatabaseError, errDB.Error())
	}

	var response dto.UserResponseDTO
	err := copier.Copy(&response, &user)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}

	return response, nil
}

func UpdateUser(userId string, userDTO dto.UpdateUserDTO) (dto.UserResponseDTO, error) {
	user, errDB := repositories.GetUserByID(userId)
	if errDB != nil {
		return dto.UserResponseDTO{}, utils.NewAppError(constants.ErrDatabaseError, errDB.Error())
	}
	errMapping := copier.Copy(&user, &userDTO)
	if errMapping != nil {
		return dto.UserResponseDTO{}, errMapping
	}
	updateUser, err := repositories.UpdateUser(user, userDTO)
	if err != nil {
		return dto.UserResponseDTO{}, err
	}

	var response dto.UserResponseDTO
	copier.Copy(&response, &updateUser)

	return response, nil
}

func DeleteUser(userId string) error {
	errDB := repositories.DeleteUser(userId)
	if errDB != nil {
		return utils.NewAppError(constants.ErrDatabaseError, errDB.Error())
	}
	return nil
}
