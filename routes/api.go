package routes

import (
	"github.com/gin-gonic/gin"
	"testgin/app/albums"
	"testgin/app/users"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r = albums.Routes(r)
	r = users.Routes(r)

	return r
}
