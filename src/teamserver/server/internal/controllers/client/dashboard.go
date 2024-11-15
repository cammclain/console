package client

import (
	"net/http"

	clientHandlers "server/internal/handlers/client"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// ClientDashboardController handles client dashboard routes
// These routes are protected by the JWT middleware
type ClientDashboardController struct {
	DB *gorm.DB
}

// NewClientDashboardController initializes a new ClientDashboardController with a database connection
func NewClientDashboardController(db *gorm.DB) *ClientDashboardController {
	return &ClientDashboardController{DB: db}
}

// GetClientDashboard handles the client dashboard route
func (c *ClientDashboardController) GetClientDashboard(ctx echo.Context) error {
	// TODO: Implement the client dashboard logic
	return ctx.JSON(http.StatusOK, map[string]string{"message": "Client dashboard"})
}

// GetClientDashboardViewCampaigns handles the client view campaigns route
func (c *ClientDashboardController) GetClientDashboardViewCampaigns(ctx echo.Context) error {
	campaigns := clientHandlers.GetCampaignsFromTeamServer(ctx)
	if campaigns == nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "No campaigns found"})
	}
	return ctx.JSON(http.StatusOK, campaigns)
}

// GetClientDashboardViewNotifications handles the client view notifications route
func (c *ClientDashboardController) GetClientDashboardViewNotifications(ctx echo.Context) error {
	notifications := clientHandlers.GetNotificationsFromTeamServer(ctx)
	// if no notifications are found, return an error
	if notifications == nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "No notifications found"})
	}
	// return the notifications
	return ctx.JSON(http.StatusOK, notifications)
}
