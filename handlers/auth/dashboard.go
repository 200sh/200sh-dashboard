package auth

import (
	"github.com/200sh/200sh-dashboard/views/dashboard"
	"github.com/labstack/echo/v4"
)

func (ah *Handler) HomeHandler(c echo.Context) error {
	return dashboard.Home(c.Path(), ah.HankoApiUrl).Render(c.Response().Writer)
}

func (ah *Handler) MonitorsHandler(c echo.Context) error {
	return dashboard.Monitor(c.Path(), ah.HankoApiUrl).Render(c.Response().Writer)
}
