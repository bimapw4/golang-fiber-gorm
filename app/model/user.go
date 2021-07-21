package model

import (
	"golang-fiber-gorm/config"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string    `json:"username"`
	Address  string    `json:"address"`
	TTL      time.Time `json:"ttl"`
}

func CreateUser(user interface{}) {
	config.DB.Create(user)
}
