package main

import (
	"net/http"

	"github.com/SE67/config"
	"github.com/SE67/controller/events"
	"github.com/SE67/controller/genders"
	"github.com/SE67/controller/order"
	"github.com/SE67/controller/paths"
	"github.com/SE67/controller/payment"
	"github.com/SE67/controller/ticket"
	"github.com/SE67/controller/timer"
	"github.com/SE67/controller/tpackage"
	"github.com/SE67/controller/users"
	"github.com/SE67/middlewares"
	"github.com/gin-gonic/gin"
)

const PORT = "8000"

func main() {

	// open connection database
	config.ConnectionDB()

	// Generate databases
	config.SetupDatabase()

	r := gin.Default()

	r.Use(CORSMiddleware())

	// Auth Route
	r.POST("/signup", users.SignUp)
	r.POST("/signin", users.SignIn)

	router := r.Group("/")
	{
		router.Use(middlewares.Authorizes())

		// User Route
		router.PUT("/user/:id", users.Update)
		router.GET("/users", users.GetAll)
		router.GET("/user/:id", users.Get)
		router.DELETE("/user/:id", users.Delete)

		// Timer Route
		router.GET("/timers", timer.GetTimers)
		router.PUT("/timers", timer.UpdateTimer)

		// Paths Route - Updated path to avoid conflict
		router.GET("/paths/:id", paths.GetPath)
		router.POST("/paths", paths.CreatePath)
		router.PUT("/paths/:id", paths.UpdatePath)
		router.DELETE("/paths/:id", paths.DeletePath)

		// Events Route - Updated path to avoid conflict
		router.GET("/events", events.GetAll)
		router.GET("/events/:id", events.GetEvent)
		router.POST("/events", events.CreateEvent)
		router.PUT("/events/:id", events.UpdateEvent)
		router.DELETE("/events/:id", events.DeleteEvent)

		// Define routes for tickets
		router.GET("/tickets", ticket.GetAll)
		router.GET("/tickets/:id", ticket.Get)
		router.POST("/tickets", ticket.CreateTicket)
		router.PUT("/tickets/:id", ticket.Update)
		router.DELETE("/tickets/:id", ticket.Delete)
		router.GET("/ticketsOrder/:id", ticket.GetTicketByOrderID)

		// Define routes for packages
		router.GET("/tpackages", tpackage.GetAll)
		router.GET("/tpackages/:id", tpackage.Get)
		router.POST("/tpackages", tpackage.CreatePackage)
		router.PUT("/tpackages/:id", tpackage.Update)
		router.DELETE("/tpackages/:id", tpackage.Delete)


		// Define routes for orders
		router.GET("/orders", order.GetAll)
		router.GET("/orders/:id", order.Get)
		router.POST("/orders", order.CreateOrder)
		router.PUT("/orders/:id", order.Update)
		router.DELETE("/orders/:id", order.Delete)

		// Define routes for payments
		router.GET("/payments", payment.GetAll)
		router.GET("/payments/:id", payment.Get)
		router.POST("/payments", payment.CreatePayment)
		router.PUT("/payments/:id", payment.Update)
		router.DELETE("/payments/:id", payment.Delete)
	}

	r.GET("/genders", genders.GetAll)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "API RUNNING... PORT: %s", PORT)
	})

	// Run the server
	r.Run("localhost:" + PORT)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
