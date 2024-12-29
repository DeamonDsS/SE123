package entity

import (
	"time"

	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model

	TotalPayment float32   `json:"total_price"`
	PaymentDate  time.Time `json:"payment_date"`
	PaymentType  string    `json:"payment_type"`

	UserID uint   `json:"user_id"`
	User   *Users `gorm:"foreignKey:UserID" json:"user"`

	OrderID uint   `json:"order_id"` // Match the Go field name
	Order   *Order `gorm:"foreignKey:OrderID" json:"order"`
}
