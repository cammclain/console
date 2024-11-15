package client

import (
	"net/http"

	database "server/internal/models"

	"github.com/labstack/echo/v4"
)

func GetCampaignsFromTeamServer(ctx echo.Context) error {
	// TODO: Implement the logic to get campaigns from the team server database
	// Make sure to check if the client has access to the campaign
	// Return the campaigns

	// database query using gorm
	var campaigns []database.Campaign

	// return the campaigns
	return ctx.JSON(http.StatusOK, campaigns)
}
