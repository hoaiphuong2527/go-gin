package repositories

import (
	"go-gin-framework/config"
	"go-gin-framework/dto"
	"go-gin-framework/models"
)

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	return users, result.Error
}

func CreateUser(user models.User) error {
	result := config.DB.Create(&user)
	return result.Error
}

func GetUserByID(id string) (models.User, error) {
	var user models.User
	err := config.DB.First(&user, "id = ?", id).Error
	return user, err
}

func UpdateUser(user models.User, userData dto.UpdateUserDTO) (models.User, error) {
	// var user models.User
	// if err := config.DB.First(&user, "id = ?", id).Error; err != nil {
	// 	return user, err
	// }

	config.DB.Model(&user).Updates(userData)
	return user, nil
}

func DeleteUser(id string) error {
	return config.DB.Delete(&models.User{}, "id = ?", id).Error
}
