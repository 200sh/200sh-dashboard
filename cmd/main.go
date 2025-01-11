package main

import (
	"context"
	"fmt"
	"github.com/200sh/200sh-dashboard/config"
	"github.com/200sh/200sh-dashboard/database"
	"github.com/200sh/200sh-dashboard/handlers"
	"github.com/200sh/200sh-dashboard/handlers/auth"
	"github.com/200sh/200sh-dashboard/hanko"
	"github.com/200sh/200sh-dashboard/internal/repository"
	middleware2 "github.com/200sh/200sh-dashboard/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

const (
	PORT = 3000
)

func main() {
	fmt.Println("⌨️Loading config")
	cfg := config.LoadConfig()

	// Database
	fmt.Println("💽Loading database")
	ctx := context.Background()
	db, err := database.New(cfg.DatabaseName)
	if err != nil {
		log.Fatalf("Error creating database: %s\n", err)
	}

	fmt.Printf("💻Starting server on http://localhost:%d\n", PORT)
	e := echo.New()

	// Pre Middlewares
	e.Pre(middleware.RemoveTrailingSlash())
	// Helper Middlewares
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Set log level
	e.Logger.SetLevel(cfg.LogLevel)

	// Setup services
	repo := repository.New(db)
	hankoClient := hanko.New(cfg.HankoApiUrl)

	// AuthMiddleware
	am := middleware2.AuthMiddleware{Hanko: &hankoClient, Ctx: ctx, Repo: repo}
	e.Use(am.IsLoggedInEnriched())

	// Setup handler
	ah := auth.Handler{Hanko: &hankoClient, Ctx: ctx, Repo: repo}

	// Setup routes
	e.Static("/static", "public")

	if cfg.IsDev {
		e.GET("/*", func(c echo.Context) error {
			response := c.Response()
			response.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
			response.Header().Set("Expires", "0")
			response.Header().Set("Pragma", "no-cache")
			response.Header().Set("Surrogate-Control", "no-store")
			return c.File(c.Param("*"))
		})
	}

	handlers.SetupRoutes(e, &am, &ah)

	// Start Server
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", PORT)))
}
