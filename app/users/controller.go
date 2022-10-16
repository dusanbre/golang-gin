package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testgin/database"
)

func Index(c *gin.Context) {
	var results []UserModel
	database.DB.Find(&results)

	c.JSON(http.StatusOK, gin.H{"data": results})
}

func Store(c *gin.Context) {
	var request CreateRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := UserModel{Name: request.Name, Email: request.Email}

	database.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
