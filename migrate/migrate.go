package main

import (
	"github.com/ryanterronez/GoJSONCRUD/initializers"
	"github.com/ryanterronez/GoJSONCRUD/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
