package tpackage

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/SE67/config"
	"github.com/SE67/entity"
)

// GetAll retrieves all Packages
func GetAll(c *gin.Context) {
	var packages []entity.Tpackage

	db := config.DB()
	results := db.Find(&packages)

	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, packages)
}

// Get retrieves a single Package by ID
func Get(c *gin.Context) {
	ID := c.Param("id")
	var tpackage entity.Tpackage

	db := config.DB()
	results := db.First(&tpackage, "id = ?", ID)

	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}

	if tpackage.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}

	c.JSON(http.StatusOK, tpackage)
}


// Update modifies an existing Package
func Update(c *gin.Context) {
	var tpackage entity.Tpackage

	PackageID := c.Param("id")

	db := config.DB()
	result := db.First(&tpackage, PackageID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ID not found"})
		return
	}

	if err := c.ShouldBindJSON(&tpackage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&tpackage)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}

func CreatePackage(c *gin.Context) {
	var tpackage entity.Tpackage

	if err := c.ShouldBindJSON(&tpackage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	db := config.DB()
	result := db.Create(&tpackage)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create package"})
		return
	}

	c.JSON(http.StatusCreated, tpackage)
}

// Delete removes a Package by ID
func Delete(c *gin.Context) {
	id := c.Param("id")

	db := config.DB()
	if tx := db.Exec("DELETE FROM tpackages WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
