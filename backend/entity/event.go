package entity


import (

   "time"

   "gorm.io/gorm"

)

type Events struct {

   gorm.Model

   EevntName string    `json:"event_name"`

   Detail  string    `json:"detail"`

   Cover     string   `gorm:"type:text" json:"cover"`

   IsPublic       uint8     `json:"ispublic"`

   Start  time.Time    `json:"start_event-"`

   End  time.Time `json:"end_event"`

   TpyeEventID  uint      `json:"type_id"`
   TpyeEvent   *TpyeEvents  `gorm:"foreignKey: TpyeEventID" `

   LocationID uint 		`json:"location_id"`
   Location		*Locations `gorm:"foreignKey: LocationID" `

   UserID uint `json:"admin_id"`
   User	*Users  `gorm:"foreignKey: UserID" `
}


