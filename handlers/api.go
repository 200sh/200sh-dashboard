package handlers

import (
	"github.com/200sh/200sh-dashboard/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func SetupApi(e *echo.Echo, cfg *config.Config) {
	a := e.Group("/api/v1")

	adm := a.Group("/admin", middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == cfg.AdminApiAuthToken, nil
	}))

	adm.GET("/monitors", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})
}
