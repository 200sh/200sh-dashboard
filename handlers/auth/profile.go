package auth

import (
	"database/sql"
	"errors"
	"github.com/200sh/200sh-dashboard/middleware"
	"github.com/200sh/200sh-dashboard/models"
	"github.com/200sh/200sh-dashboard/views/dashboard"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) ProfileHandler(c echo.Context) error {
	user := c.Get(middleware.UserIDKey).(*models.User)
	return renderProfilePage(c, h.hanko.HankoApiUrl, user, "")
}

type ProfileForm struct {
	Name string `form:"name"`
}

func (h *Handler) ProfileFormHandler(c echo.Context) error {
	user := c.Get(middleware.UserIDKey).(*models.User)

	var form ProfileForm
	if err := c.Bind(&form); err != nil {
		return c.String(http.StatusBadRequest, "Invalid form data")
	}

	// Update temp values
	user.Name = form.Name

	if err := h.userService.Update(user); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.String(http.StatusNotFound, "User not found")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusSeeOther, "/dashboard/profile")
}

func renderProfilePage(c echo.Context, hankoUrl string, user *models.User, errorMsg string) error {
	return dashboard.Profile(
		c.Request().URL.Path,
		hankoUrl,
		user,
		errorMsg,
	).Render(c.Response().Writer)
}
