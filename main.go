package main

import (
	"fmt"
	"github.com/200sh/200sh-dashboard/pages"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	fmt.Println("Loading templates from ./templates")
	t := pages.LoadPageTemplates()

	fmt.Println("Starting server on http://localhost:8080")
	e := echo.New()
	// Setup root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Setup the renderer
	e.Renderer = t

	// Setup Static files from the 'public' folder
	e.Static("/", "public")

	// Define routes
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})
	e.GET("/privacy", func(c echo.Context) error {
		return c.Render(http.StatusOK, "privacy.html", nil)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
