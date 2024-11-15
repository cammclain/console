package client

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// ClientAuthController handles client authentication routes
type ClientAuthController struct {
	DB *gorm.DB
}

// NewClientAuthController initializes a new ClientAuthController with a database connection
func NewClientAuthController(db *gorm.DB) *ClientAuthController {
	return &ClientAuthController{DB: db}
}

// ClientLogin handles the login route for the desktop client
func (c *ClientAuthController) ClientLogin(ctx echo.Context) error {
	// TODO: Implement the login logic
	// Example: check credentials, generate a token, return it

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Login successful",
		"token":   "example_token_here",
	})
}

// ClientRequestAnAccount handles the request for an account route for the desktop client
func (c *ClientAuthController) ClientRequestAnAccount(ctx echo.Context) error {
	// TODO: Implement the request for an account logic
	// Sends a notification to the team server to create an account for the client
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Request for an account sent",
	})
}
