package main

import (
	"go-gin-framework/config"
	"go-gin-framework/routes"
)

func main() {
	config.ConnectDatabase()

	r := routes.SetupRouter()
	r.Run(":8080")
}
