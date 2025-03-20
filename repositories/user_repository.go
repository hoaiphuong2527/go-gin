package repositories

import (
	"go-gin-framework/config"
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
