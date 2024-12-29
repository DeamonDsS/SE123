package ticket

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/SE67/config"
	"github.com/SE67/entity"
)

// GetAll retrieves all Tickets
func GetAll(c *gin.Context) {
	var tickets []entity.Ticket

	db := config.DB()
	results := db.Preload("Code").Preload("Package").Preload("Order").Find(&tickets)

	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, tickets)
}

// Get retrieves a single Ticket by ID
func Get(c *gin.Context) {
	ID := c.Param("id")
	var ticket entity.Ticket

	db := config.DB()
	results := db.Preload("Code").Preload("Package").Preload("Order").First(&ticket, ID)

	if results.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, ticket)
}

func GetTicketByOrderID(c *gin.Context) {
    orderID := c.Param("orderID")
    var tickets []entity.Ticket

    db := config.DB()
    results := db.Preload("Code").Preload("Package").Preload("Order").
        Where("order_id = ?", orderID).Find(&tickets)

    if results.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": results.Error.Error()})
        return
    }

    if len(tickets) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"message": "No tickets found for this order ID"})
        return
    }

    c.JSON(http.StatusOK, tickets)
}

// Update modifies an existing Ticket
func Update(c *gin.Context) {
	var updates map[string]interface{}
	TicketID := c.Param("id")

	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	db := config.DB()
	result := db.Model(&entity.Ticket{}).Where("id = ?", TicketID).Updates(updates)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to update ticket: " + result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Updated successfully"})
}

// CreateTicket creates a new Ticket
func CreateTicket(c *gin.Context) {
	var ticket entity.Ticket

	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	db := config.DB()
	result := db.Create(&ticket)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create ticket: " + result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, ticket)
}

// Delete removes a Ticket by ID
func Delete(c *gin.Context) {
	id := c.Param("id")

	db := config.DB()
	result := db.Delete(&entity.Ticket{}, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Deleted successfully"})
}
