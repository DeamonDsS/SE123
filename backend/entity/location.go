package entity


import "gorm.io/gorm"



type Locations struct{
	gorm.Model
	Name string `json:"location_name"`
 }