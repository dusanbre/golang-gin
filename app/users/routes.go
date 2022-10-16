package users

import (
	"github.com/gin-gonic/gin"
	"testgin/app/middleware"
)

func Routes(r *gin.Engine) *gin.Engine {
	r.GET("/users", Index)
	r.POST("/users", Store)
	r.POST("/users/login", Login)
	secured := r.Group("/secured").Use(middleware.AuthJWT())
	{
		secured.GET("/ping", Ping)
	}
	return r
}
