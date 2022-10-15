package albums

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	var data = GetAlbums()
	c.IndentedJSON(http.StatusOK, data)
}
