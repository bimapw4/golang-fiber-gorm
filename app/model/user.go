package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string    `json:"username"`
	Address  string    `json:"address"`
	TTL      time.Time `json:"ttl"`
}
