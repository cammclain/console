// src/teamserver/server/main.go
package main

import (
	"log"
	"server/internal/database"
	"server/internal/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Initialize Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize Database
	db := database.Connect() // Assuming `database.Connect` returns *gorm.DB
	if db == nil {
		log.Fatalf("Failed to connect to database")
	}

	// Register routes
	routes.RegisterClientRoutes(e, db)
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
