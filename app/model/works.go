package model

import (
	"golang-fiber-gorm/config"

	"gorm.io/gorm"
)

type Works struct {
	gorm.Model
	Work        string `json:"work"`
	Description string `json:"description"`
}

func CreateWorks(Works interface{}) *gorm.DB {
	return config.DB.Create(Works)
}

func UpdateWorksById(id interface{}, data interface{}) *gorm.DB {
	return config.DB.Model(&Works{}).Where("id = ?", id).Updates(data)
}

func GetWorksList(resp interface{}) *gorm.DB {
	return config.DB.Model(&Works{}).Find(resp)
}

func DeleteWorks(Condition interface{}) *gorm.DB {
	return config.DB.Where(Condition).Delete(&Works{})
}
