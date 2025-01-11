package auth

import (
	"github.com/200sh/200sh-dashboard/internal/repository"
	"github.com/200sh/200sh-dashboard/middleware"
	"github.com/200sh/200sh-dashboard/views/dashboard"
	"github.com/labstack/echo/v4"
	log2 "github.com/labstack/gommon/log"
	"net/http"
	"net/url"
	"strconv"
)

func (h *Handler) HomeHandler(c echo.Context) error {
	user := c.Get(middleware.UserIDKey).(*repository.User)
	if user == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	return dashboard.Home(c.Path(), h.Hanko.HankoApiUrl, user).Render(c.Response().Writer)
}

func (h *Handler) MonitorsHandler(c echo.Context) error {
	user := c.Get(middleware.UserIDKey).(*repository.User)
	if user == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	monitors, err := h.Repo.GetMonitorsByUserID(h.Ctx, user.ID)
	if err != nil {
		return err
	}

	return dashboard.Monitor(c.Path(), h.Hanko.HankoApiUrl, user, monitors).Render(c.Response().Writer)
}

func (h *Handler) ViewMonitorHandler(c echo.Context) error {
	user := c.Get(middleware.UserIDKey).(*repository.User)
	if user == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/dashboard/monitors")
	}

	monitor, err := h.Repo.GetMonitorByUserID(h.Ctx, repository.GetMonitorByUserIDParams{
		UserID: user.ID,
		ID:     int64(id),
	})
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/dashboard/monitors")
	}

	return dashboard.ViewMonitor(c.Path(), h.Hanko.HankoApiUrl, user, &monitor).Render(c.Response().Writer)
}

func (h *Handler) NewMonitorHandler(c echo.Context) error {
	user := c.Get(middleware.UserIDKey).(*repository.User)
	if user == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	return dashboard.NewMonitor(h.Hanko.HankoApiUrl, user).Render(c.Response().Writer)
}

type NewMonitorForm struct {
	Url string `form:"monitor-url"`
}

func (h *Handler) NewMonitorFormHandler(c echo.Context) error {
	var nmf NewMonitorForm
	err := c.Bind(&nmf)
	if err != nil {
		return c.String(http.StatusBadRequest, "Not valid form")
	}

	user := c.Get(middleware.UserIDKey).(*repository.User)
	if user == nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	log2.Info("Got request form: ", nmf)

	log2.Info("From user: ", user)

	mUrl, err := url.Parse(nmf.Url)
	if err != nil {
		// TODO: should return to page with some error message
		return err
	}

	// Create the new monitor object
	_, err = h.Repo.CreateMonitor(h.Ctx, repository.CreateMonitorParams{
		UserID: user.ID,
		Url:    mUrl.String(),
	})
	if err != nil {
		// TODO: Should add some error message to the form
		return c.Redirect(http.StatusSeeOther, "/dashboard/monitors/new")
	}

	// TODO: Should have some some success toast?
	return c.Redirect(http.StatusSeeOther, "/dashboard/monitors")
}
