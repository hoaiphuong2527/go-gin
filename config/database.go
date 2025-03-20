package config

import (
	"fmt"
	"go-gin-framework/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/gin?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Cannot connect to MySQL:", err)
	}

	fmt.Println("Connected to MySQL")
	DB = database

	// Tạo bảng nếu chưa có
	database.AutoMigrate(&models.User{})
}
