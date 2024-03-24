package main

import (
	"crud/initializers"
	"crud/models"
)

func init() {
	initializers.LoadEnvironmentVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
