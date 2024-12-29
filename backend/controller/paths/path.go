package paths


import (

   "net/http"


   "github.com/SE67/config"

    "github.com/SE67/entity"

   "github.com/gin-gonic/gin"

)



// GetPath - ดึงข้อมูล Path โดยระบุ ID
func GetPath(c *gin.Context) {
	var path entity.Path
	db := config.DB()
	id := c.Param("id") // รับ ID จากพารามิเตอร์ URL

	if err := db.Preload("Timer").Preload("Location").Preload("User").Where("id = ?", id).First(&path).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Path not found"})
		return
	}

	c.JSON(http.StatusOK, path)
}

// UpdatePath - อัปเดตข้อมูล Path ที่ระบุโดย ID
func UpdatePath(c *gin.Context) {
	id := c.Param("id") // รับ ID จากพารามิเตอร์ URL
	var path entity.Path
	db := config.DB()

	// ตรวจสอบว่ามี Path อยู่ในระบบหรือไม่
	if err := db.First(&path, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Path not found"})
		return
	}

	var updateData entity.Path
	// อ่านข้อมูล JSON จาก Request Body
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// อัปเดตฟิลด์ที่ต้องการ
	path.FirstTime = updateData.FirstTime
	path.NextTime = updateData.NextTime
	path.TimeToNext = updateData.TimeToNext
	path.TimerID = updateData.TimerID
	path.LocationID = updateData.LocationID
	path.UserID = updateData.UserID

	// บันทึกการเปลี่ยนแปลงลงฐานข้อมูล
	if err := db.Save(&path).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, path)
}

// CreatePath - สร้าง Path ใหม่
func CreatePath(c *gin.Context) {
	var newPath entity.Path
	db := config.DB()

	// อ่านข้อมูล JSON จาก Request Body
	if err := c.ShouldBindJSON(&newPath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// สร้าง Path ใหม่ในฐานข้อมูล
	if err := db.Create(&newPath).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newPath)
}

// DeletePath - ลบ Path ที่ระบุโดย ID
func DeletePath(c *gin.Context) {
	id := c.Param("id") // รับ ID จากพารามิเตอร์ URL
	db := config.DB()

	// ลบ Path จากฐานข้อมูล
	if err := db.Delete(&entity.Path{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Path deleted successfully"})
}
