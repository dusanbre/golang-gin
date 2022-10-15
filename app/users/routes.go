package users

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) *gin.Engine {
	r.GET("/users", Index)
	return r
}
