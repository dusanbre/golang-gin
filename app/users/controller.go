package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testgin/database"
)

func Index(c *gin.Context) {
	db := database.GetDB()

	result := map[string]interface{}{}
	db.Model(&UserModel{}).First(&result)

	c.JSON(http.StatusOK, result)
}
