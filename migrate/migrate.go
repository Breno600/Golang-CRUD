package main

import (
	"github.com/breno600/go-crud/initializers"
	"github.com/breno600/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}