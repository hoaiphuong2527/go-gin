package main

import (
	"go-gin-framework/config"
	"go-gin-framework/routes"
	"log"
)

func main() {
	config.ConnectDatabase()
	config.LoadEnv()
	port := config.GetEnv("SERVER_PORT", "8080")
	r := routes.SetupRouter()
	r.Run(":" + port)
	err := r.Run(":" + port)
	log.Println("🚀 Server is running on port:", port)
	if err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
