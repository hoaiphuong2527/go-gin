package services

import (
	"go-gin-framework/constants"
	"go-gin-framework/dto"
	"go-gin-framework/models"
	"go-gin-framework/repositories"
	"go-gin-framework/utils"
	"strconv"

	"github.com/gin-gonic/gin"
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

func GetAllUsers(context *gin.Context) (dto.PaginatedResponse, error) {
	page, _ := strconv.Atoi(context.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	users, total, errGetList := repositories.GetUsers(context, page, pageSize)
	if errGetList != nil {
		return dto.PaginatedResponse{}, utils.NewAppError(constants.ErrDatabaseError, errGetList.Error())
	}
	var mapDto []dto.UserResponseDTO
	copier.Copy(&mapDto, &users)

	response := dto.PaginatedResponse{
		Data:       mapDto,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int((total + int64(pageSize) - 1) / int64(pageSize)),
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
