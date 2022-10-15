package albums

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) *gin.Engine {
	r.GET("/albums", Index)
	return r
}
