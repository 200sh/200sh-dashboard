package main

import (
	"fmt"
	"github.com/200sh/200sh-dashboard/config"
	"github.com/200sh/200sh-dashboard/handlers"
	"github.com/200sh/200sh-dashboard/handlers/auth"
	middleware2 "github.com/200sh/200sh-dashboard/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("⌨️Loading config")
	cfg := config.LoadConfig()

	fmt.Println("💻Starting server on http://localhost:8080")
	e := echo.New()

	// Pre Middlewares
	e.Pre(middleware.RemoveTrailingSlash())
	// Helper Middlewares
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Set log level
	e.Logger.SetLevel(cfg.LogLevel)

	// AuthMiddleware
	am := middleware2.AuthMiddleware{HankoApiUrl: cfg.HankoApiUrl}
	e.Use(am.IsLoggedInEnriched())

	// Setup handler
	ah := auth.Handler{HankoApiUrl: cfg.HankoApiUrl}

	// Setup routes
	e.Static("/static", "public")
	handlers.SetupRoutes(e, &am, &ah)

	// Start Server
	e.Logger.Fatal(e.Start(":8080"))
}
