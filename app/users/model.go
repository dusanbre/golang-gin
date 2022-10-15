package users

import "gorm.io/gorm"

type UserModel struct {
	Name  string
	Email *string
	gorm.Model
}
