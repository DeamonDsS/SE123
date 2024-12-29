package timer

import (

	"net/http"
 
 
	"github.com/gin-gonic/gin"
 
 
	 "github.com/SE67/config"
 
	"github.com/SE67/entity"
 
 )
 



// GetTimers - ดึงข้อมูล Timers 
func GetTimers(c *gin.Context) {
	var timer entity.Timers
	db := config.DB()
	if err := db.Where("ID = ?", 1).First(&timer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, timer)
}

// UpdateTimer - อัปเดตข้อมูล Timer ที่ระบุโดย ID
func UpdateTimer(c *gin.Context) {
	id := c.Param("1")
	var timer entity.Timers
	db := config.DB()
	if err := db.First(&timer, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Timer not found"})
		return
	}

	var updateData entity.Timers
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	timer.Start = updateData.Start
	timer.Stop = updateData.Stop
	timer.Freq_mins = updateData.Freq_mins

	if err := db.Save(&timer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, timer)
}



