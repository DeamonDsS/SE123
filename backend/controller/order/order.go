package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/SE67/config"
	"github.com/SE67/entity"
)

// GetAll retrieves all Orders
func GetAll(c *gin.Context) {
	var orders []entity.Order

	db := config.DB()
	results := db.Preload("User").Find(&orders)

	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// Get retrieves a single Order by ID
func Get(c *gin.Context) {
	ID := c.Param("id")
	var order entity.Order

	db := config.DB()
	results := db.Preload("User").First(&order, ID)

	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}

	if order.ID == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}

	c.JSON(http.StatusOK, order)
}

// Update modifies an existing Order
func Update(c *gin.Context) {
	var order entity.Order

	OrderID := c.Param("id")

	db := config.DB()
	result := db.First(&order, OrderID)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ID not found"})
		return
	}

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	result = db.Save(&order)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}

func CreateOrder(c *gin.Context) {
	var order entity.Order

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request, unable to map payload"})
		return
	}

	db := config.DB()
	result := db.Create(&order)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	c.JSON(http.StatusCreated, order)
}

// Delete removes an Order by ID
func Delete(c *gin.Context) {
	id := c.Param("id")

	db := config.DB()
	if tx := db.Exec("DELETE FROM orders WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
