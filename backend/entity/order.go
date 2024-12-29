package entity

import (
	// "time"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model

	TotalOrder string    `json:"t_order"`
	
	UserID uint   `json:"user_id"`
	User   *Users `gorm:"foreignKey:UserID" json:"user"` // Correct foreign key reference
}
