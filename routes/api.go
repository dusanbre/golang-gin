package routes

import (
	"testgin/app/albums"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r = albums.Routes(r)

	return r
}
