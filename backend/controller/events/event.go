package events


import (

	"net/http"
 
	"github.com/SE67/config"
 
	 "github.com/SE67/entity"
 
	"github.com/gin-gonic/gin"
 
 )

// GetEvent - ดึงข้อมูล Event โดยระบุ ID

func GetAll(c *gin.Context) {


	var event []entity.Events

 
 
	db := config.DB()
 
	results := db.Find(&event)
 
	if results.Error != nil {
 
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
 
		return
 
	}
 
	c.JSON(http.StatusOK, event)
 
 
 }
func GetEvent(c *gin.Context) {
	var event entity.Events
	db := config.DB()
	id := c.Param("id") // รับ ID จากพารามิเตอร์ URL

	if err := db.Where("id = ?", id).First(&event).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	c.JSON(http.StatusOK, event)
}

// UpdateEvent - อัปเดตข้อมูล Event ที่ระบุโดย ID
func UpdateEvent(c *gin.Context) {
	id := c.Param("id") // รับ ID จากพารามิเตอร์ URL
	var event entity.Events
	db := config.DB()

	// ตรวจสอบว่ามี Event อยู่ในระบบหรือไม่
	if err := db.First(&event, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	var updateData entity.Events
	// อ่านข้อมูล JSON จาก Request Body
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// อัปเดตฟิลด์ที่ต้องการ
	event.EevntName = updateData.EevntName
	event.Detail = updateData.Detail
	event.Cover = updateData.Cover
	event.IsPublic = updateData.IsPublic
	event.Start = updateData.Start
	event.End = updateData.End
	event.TpyeEventID = updateData.TpyeEventID
	event.LocationID = updateData.LocationID
	event.UserID = updateData.UserID

	// บันทึกการเปลี่ยนแปลงลงฐานข้อมูล
	if err := db.Save(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, event)
}

// CreateEvent - สร้าง Event ใหม่
func CreateEvent(c *gin.Context) {
	var newEvent entity.Events
	db := config.DB()

	// อ่านข้อมูล JSON จาก Request Body
	if err := c.ShouldBindJSON(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// สร้าง Event ใหม่ในฐานข้อมูล
	if err := db.Create(&newEvent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newEvent)
}

// DeleteEvent - ลบ Event ที่ระบุโดย ID
func DeleteEvent(c *gin.Context) {
	id := c.Param("id") // รับ ID จากพารามิเตอร์ URL
	db := config.DB()

	// ลบ Event จากฐานข้อมูล
	if err := db.Delete(&entity.Events{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
