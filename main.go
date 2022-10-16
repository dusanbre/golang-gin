package main

import (
	"testgin/app/users"
	"testgin/database"
	"testgin/routes"
)

func main() {
	r := routes.SetupRouter()

	db := database.Init()

	db.AutoMigrate(&users.UserModel{})

	r.Run(":8000")
}
