package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"testgin/app/auth"
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

	user := UserModel{Name: request.Name, Email: request.Email, Username: request.Username, Password: request.Password}

	if err := user.HashPassword(request.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	record := database.DB.Create(&user)

	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func Login(context *gin.Context) {
	var request LoginRequest
	var user UserModel

	// Construct request struct
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	// Check if user exists
	record := database.DB.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}

	// Check password
	credentialError := user.CheckPassword(request.Password)

	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}

	// Generate JWT token
	tokenString, err := auth.GenerateJWT(user.Email, user.Username)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Secure route"})
}
