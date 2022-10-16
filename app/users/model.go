package users

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserModel struct {
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Username string `json:"username" gorm:"unique"`
	gorm.Model
}

func (user *UserModel) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *UserModel) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))

	if err != nil {
		return err
	}

	return nil
}
