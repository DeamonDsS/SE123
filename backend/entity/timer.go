package entity


import "gorm.io/gorm"


type Timers struct {
	gorm.Model
	Start string    `json:"strat"`
	Stop string    `json:"stop"`
	Freq_mins int  `json:"freq_mins"`
 }
