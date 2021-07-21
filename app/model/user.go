package model

import (
	"golang-fiber-gorm/config"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Address  string `json:"address"`
	Email    string `json:"email"`
}

func CreateUser(user interface{}) {
	config.DB.Create(user)
}

func UpdateUserById(id interface{}, data interface{}) {
	config.DB.Model(&User{}).Where("id = ?", id).Updates(data)
}
