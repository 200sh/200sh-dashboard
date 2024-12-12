package auth

import (
	"github.com/200sh/200sh-dashboard/middleware"
	"github.com/200sh/200sh-dashboard/views/auth"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	HankoApiUrl string
}

func (ah *Handler) LoginPageHandler(c echo.Context) error {
	isLoggedIn := c.Get(middleware.IsLoggedInKey).(bool)
	if isLoggedIn {
		return c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
	}
	return auth.Login(ah.HankoApiUrl).Render(c.Response().Writer)
}
