package client

import (
	"net/http"
	clientHandlers "server/internal/handlers/client"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// ClientCampaignsController handles client campaigns routes
// These routes are protected by the JWT middleware
type ClientCampaignsController struct {
	DB *gorm.DB
}

// NewClientCampaignsController initializes a new ClientCampaignsController with a database connection
func NewClientCampaignsController(db *gorm.DB) *ClientCampaignsController {
	return &ClientCampaignsController{DB: db}
}

// CampaignsOverviewView returns all campaigns from the team server that the client has access to
func (c *ClientCampaignsController) CampaignsOverviewView(ctx echo.Context) error {
	// TODO: Implement the logic to get campaigns from the team server
	// Make sure to check if the client has access to the campaign
	// Return the campaigns
	return clientHandlers.GetCampaignsFromTeamServer(ctx)
}

// ListCampaignsView returns all campaigns from the team server that the client has access to
func (c *ClientCampaignsController) ListCampaignsView(ctx echo.Context) error {
	// validate the client has access to the campaigns
	// TODO: Implement the logic to validate the client has access to the campaigns

	// query the campaigns from the team server database
	campaigns := clientHandlers.GetCampaignsFromTeamServer(ctx)
	if campaigns == nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "No campaigns found"})
	}

	return ctx.JSON(http.StatusOK, campaigns)
}
