package handlers

import (
	"github.com/200sh/200sh-dashboard/middleware"
	"github.com/200sh/200sh-dashboard/views/landing"
	"github.com/labstack/echo/v4"
)

func LandingPageHandler(c echo.Context) error {
	isLoggedIn := c.Get(middleware.IsLoggedInKey).(bool)
	return landing.Home(isLoggedIn).Render(c.Response().Writer)
}

func NotFoundPageHandler(c echo.Context) error {
	isLoggedIn := c.Get(middleware.IsLoggedInKey).(bool)
	return landing.NotFound(isLoggedIn).Render(c.Response().Writer)
}
