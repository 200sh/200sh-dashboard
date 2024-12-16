package auth

import (
	"github.com/200sh/200sh-dashboard/middleware"
	"github.com/200sh/200sh-dashboard/models"
	"github.com/200sh/200sh-dashboard/views/dashboard"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) HomeHandler(c echo.Context) error {
	user := c.Get(middleware.UserIDKey).(*models.User)
	if user == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	return dashboard.Home(c.Path(), h.Hanko.HankoApiUrl, user).Render(c.Response().Writer)
}

func (h *Handler) MonitorsHandler(c echo.Context) error {
	user := c.Get(middleware.UserIDKey).(*models.User)
	if user == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	return dashboard.Monitor(c.Path(), h.Hanko.HankoApiUrl, user).Render(c.Response().Writer)
}
