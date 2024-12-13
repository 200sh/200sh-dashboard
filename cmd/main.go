package main

import (
	"fmt"
	"github.com/200sh/200sh-dashboard/config"
	"github.com/200sh/200sh-dashboard/database"
	"github.com/200sh/200sh-dashboard/handlers"
	"github.com/200sh/200sh-dashboard/handlers/auth"
	"github.com/200sh/200sh-dashboard/hanko"
	middleware2 "github.com/200sh/200sh-dashboard/middleware"
	"github.com/200sh/200sh-dashboard/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	fmt.Println("⌨️Loading config")
	cfg := config.LoadConfig()

	// Database
	fmt.Println("💽Loading database")
	db, err := database.New(cfg.DatabaseName)
	if err != nil {
		log.Fatalf("Error creating database: %s\n", err)
	}

	fmt.Println("💻Starting server on http://localhost:8080")
	e := echo.New()

	// Pre Middlewares
	e.Pre(middleware.RemoveTrailingSlash())
	// Helper Middlewares
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Set log level
	e.Logger.SetLevel(cfg.LogLevel)

	// Setup services
	us := models.NewUserService(db)
	hankoClient := hanko.New(cfg.HankoApiUrl)

	// AuthMiddleware
	am := middleware2.AuthMiddleware{Hanko: &hankoClient, UserService: us}
	e.Use(am.IsLoggedInEnriched())

	// Setup handler
	ah := auth.Handler{Hanko: &hankoClient, UserService: us}

	// Setup routes
	e.Static("/static", "public")
	handlers.SetupRoutes(e, &am, &ah)

	// Start Server
	e.Logger.Fatal(e.Start(":8080"))
}
