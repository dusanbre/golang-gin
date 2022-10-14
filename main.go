package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		data := map[string]interface{}{
			"lang":      "GOLANG",
			"framework": "GIN",
		}

		c.AsciiJSON(http.StatusOK, data)
	})

	r.Run(":8000")
}
