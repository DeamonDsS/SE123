package entity


import (

   "gorm.io/gorm"

)

type Tpackage struct {

   gorm.Model

   Name string    `json:"t_name" gorm:"unique"`
 
   Type  string    `json:"t_type" gorm:"default:'standard'"`

   Price float32    `json:"t_price"`
   
   Zone string     `json:"t_zone"`

   Description string `json:"t_des"`

}