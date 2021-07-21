package model

import (
	"gorm.io/gorm"
)

type Works struct {
	gorm.Model
	Work        string `json:"work"`
	Description string `json:"description"`
}
