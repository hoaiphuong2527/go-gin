package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	// "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	LoadEnv()

	dbDriver := GetEnv("DB_DRIVER", "mysql")
	dbUser := GetEnv("DB_USER", "root")
	dbPassword := GetEnv("DB_PASSWORD", "")
	dbHost := GetEnv("DB_HOST", "localhost")
	dbPort := GetEnv("DB_PORT", "3306")
	dbName := GetEnv("DB_NAME", "mydatabase")
	var err error
	if dbDriver == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			dbUser, dbPassword, dbHost, dbPort, dbName)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	} else if dbDriver == "postgres" {
		// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		// 	dbHost, dbUser, dbPassword, dbName, dbPort)
		// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}

	if err != nil {
		log.Fatal("🚀 Failed to connect to database:", err)
	}

	log.Println("✅ Connected to the database successfully!")
}
