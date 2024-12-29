package entity

import (
	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model

	OwnerName string    `json:"owner_name" valid:"required~owner_name is required"`
	Phone	string 	    `json:"phone" valid:"required~phone is required"`
	
	CodeID    uint      `json:"code_id"`
    Code      *Code     `gorm:"foreignKey:CodeID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"code"`

	PackageID uint      `json:"tpackage_id"`
	Package   *Tpackage `gorm:"foreignKey:PackageID" json:"package"` // Correct foreign key reference

	OrderID uint   `json:"order_id"` // Match the Go field name
	Order   *Order `gorm:"foreignKey:OrderID" json:"order"`
}
