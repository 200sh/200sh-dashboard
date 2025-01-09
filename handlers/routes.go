package handlers

import (
	"github.com/200sh/200sh-dashboard/handlers/auth"
	"github.com/200sh/200sh-dashboard/middleware"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, am *middleware.AuthMiddleware, ah *auth.Handler) {
	e.GET("/", LandingPageHandler)
	e.GET("/login", ah.LoginPageHandler)
	e.GET("/user/setup", ah.UserSetupPage)
	e.POST("/user/setup", ah.UserSetupForm)

	// Auth Required routes
	aq := e.Group("/dashboard", am.AuthRequired())
	aq.GET("", ah.HomeHandler)
	aq.GET("/monitors", ah.MonitorsHandler)
	aq.GET("/monitors/:id", ah.ViewMonitorHandler)
	aq.GET("/monitors/new", ah.NewMonitorHandler)
	aq.POST("/monitors/new", ah.NewMonitorFormHandler)

	// Route not found fallback
	e.GET("/*", NotFoundPageHandler)
}
