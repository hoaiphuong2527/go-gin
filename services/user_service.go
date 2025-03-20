package services

import (
	"go-gin-framework/models"
	"go-gin-framework/repositories"
)

func FetchAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

func RegisterUser(user models.User) error {
	return repositories.CreateUser(user)
}
