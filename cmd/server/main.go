package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/yyoosshh/data-steward/config"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize Echo instance
	e := echo.New()

	// TODO: Setup routes, middleware, etc.

	// Start server
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
