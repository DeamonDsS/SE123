package entity


import "gorm.io/gorm"



type Path struct {

	gorm.Model
 
	FirstTime string    `json:"first_time"`
	NextTime  string	`json:"next_time"`
	TimeToNext int      `json:"time_to_next"`
	TimerID  uint      `json:"type_id"`
	Timer   *Timers    `gorm:"foreignKey: TimerID" `
	LocationID uint 		`json:"location_id"`
	Location	*Locations `gorm:"foreignKey: LocationID" `
    UserID   uint `json:"admin_id"`
    User	*Users  `gorm:"foreignKey: UserID" `
	
 }
