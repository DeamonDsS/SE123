package entity

import (
	"gorm.io/gorm"
)

type Code struct {
	gorm.Model

	Code string    `json:"code"`
	
}
