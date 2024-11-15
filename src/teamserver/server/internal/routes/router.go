package routes

import (
	"server/internal/controllers/client"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterClientRoutes(e *echo.Echo, db *gorm.DB) {
	// create the client auth controller
	clientAuthController := client.NewClientAuthController(db)

	// group the routes
	clientRoutes := e.Group("/client")

	// register the routes

	clientRoutes.POST("/login", clientAuthController.ClientLogin)
	clientRoutes.POST("/request-an-account", clientAuthController.ClientRequestAnAccount)
}
