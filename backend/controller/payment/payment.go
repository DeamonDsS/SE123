package payment

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/SE67/config"
	"github.com/SE67/entity"
)

// GetAll retrieves all Payments
func GetAll(c *gin.Context) {
	var payments []entity.Payment

	db := config.DB()
	results := db.Preload("User").Preload("Order").Find(&payments)

	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}

// Get retrieves a single Payment by ID
func Get(c *gin.Context) {
	ID := c.Param("id")
	var payment entity.Payment

	db := config.DB()
	results := db.Preload("User").Preload("Order").First(&payment, ID)

	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}

	if payment.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}

	c.JSON(http.StatusOK, payment)
}

// Update modifies an existing Payment
func Update(c *gin.Context) {
	var payment entity.Payment

	PaymentID := c.Param("id")

	db := config.DB()
	result := db.First(&payment, PaymentID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ID not found"})
		return
	}

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&payment)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}

func CreatePayment(c *gin.Context) {
	var payment entity.Payment

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	db := config.DB()
	result := db.Create(&payment)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
		return
	}

	c.JSON(http.StatusCreated, payment)
}
// Delete removes a Payment by ID
func Delete(c *gin.Context) {
	id := c.Param("id")

	db := config.DB()
	if tx := db.Exec("DELETE FROM payments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
