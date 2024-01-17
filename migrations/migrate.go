package main

import (
	"github.com/yoma/go-crud/initializers"
	"github.com/yoma/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}

// running this file should create our table and database