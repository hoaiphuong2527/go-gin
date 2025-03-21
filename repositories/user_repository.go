package repositories

import (
	"go-gin-framework/config"
	"go-gin-framework/dto"
	"go-gin-framework/models"

	"github.com/gin-gonic/gin"
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
	config.DB.Model(&user).Updates(userData)
	return user, nil
}

func DeleteUser(id string) error {
	return config.DB.Delete(&models.User{}, "id = ?", id).Error
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := config.DB.First(&user, "email = ?", email).Error
	return user, err
}

func GetUsers(context *gin.Context, page int, pageSize int) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	query := config.DB.Model(&models.User{})

	if name := context.Query("name"); name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}
	if email := context.Query("email"); email != "" {
		query = query.Where("email LIKE ?", "%"+email+"%")
	}

	query.Count(&total)

	offset := (page - 1) * pageSize
	err := query.Offset(offset).Limit(pageSize).Find(&users).Error

	return users, total, err
}
