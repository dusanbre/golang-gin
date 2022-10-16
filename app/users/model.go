package users

import "gorm.io/gorm"

type UserModel struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	gorm.Model
}
