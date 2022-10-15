package albums

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var data = GetAlbums()
	c.IndentedJSON(http.StatusOK, data)
}
