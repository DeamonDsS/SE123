package entity


import "gorm.io/gorm"



type TpyeEvents struct {

	gorm.Model
 
	TpyeName string    `json:"type_name"`
 
 }
