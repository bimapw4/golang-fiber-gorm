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
	Work_id  int    `json:"work_id"`
}

func CreateUser(user interface{}) {
	config.DB.Create(user)
}

func UpdateUserById(id interface{}, data interface{}) {
	config.DB.Model(&User{}).Where("id = ?", id).Updates(data)
}

func GetUserList(resp interface{}) *gorm.DB {
	return config.DB.Model(&User{}).Find(resp)
}

func DeleteUser(Condition interface{}) *gorm.DB {
	return config.DB.Where(Condition).Delete(&User{})
}
