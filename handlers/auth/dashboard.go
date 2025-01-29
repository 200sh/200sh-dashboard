package auth

import (
	"github.com/200sh/200sh-dashboard/middleware"
	"github.com/200sh/200sh-dashboard/models"
	"github.com/200sh/200sh-dashboard/views/dashboard"
	"github.com/labstack/echo/v4"
	log2 "github.com/labstack/gommon/log"
	"net/http"
	"net/url"
	"strconv"
)

func (h *Handler) HomeHandler(c echo.Context) error {
	user := c.Get(middleware.UserIDKey).(*models.User)
	if user == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	return dashboard.Home(c.Path(), h.hanko.HankoApiUrl, user).Render(c.Response().Writer)
}

func (h *Handler) MonitorsHandler(c echo.Context) error {
	user := c.Get(middleware.UserIDKey).(*models.User)
	if user == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	monitors, err := h.monitorService.ListByUser(user.Id)
	if err != nil {
		return err
	}

	return dashboard.Monitor(c.Path(), h.hanko.HankoApiUrl, user, monitors).Render(c.Response().Writer)
}

func (h *Handler) ViewMonitorHandler(c echo.Context) error {
	user := c.Get(middleware.UserIDKey).(*models.User)
	if user == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/dashboard/monitors")
	}

	monitor, err := h.monitorService.GetByIDAndUser(id, user.Id)
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/dashboard/monitors")
	}

	return dashboard.ViewMonitor(c.Path(), h.hanko.HankoApiUrl, user, monitor).Render(c.Response().Writer)
}

func (h *Handler) NewMonitorHandler(c echo.Context) error {
	user := c.Get(middleware.UserIDKey).(*models.User)
	if user == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	return dashboard.NewMonitor(h.hanko.HankoApiUrl, user).Render(c.Response().Writer)
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

	user := c.Get(middleware.UserIDKey).(*models.User)
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
	m := &models.Monitor{
		UserId: user.Id,
		Url:    mUrl.String(),
	}
	err = h.monitorService.Create(m)
	if err != nil {
		// TODO: Should add some error message to the form
		return c.Redirect(http.StatusSeeOther, "/dashboard/monitors/new")
	}

	// TODO: Should have some some success toast?
	return c.Redirect(http.StatusSeeOther, "/dashboard/monitors")
}

func (h *Handler) DeleteMonitorHandler(c echo.Context) error {
	user := c.Get(middleware.UserIDKey).(*models.User)
	if user == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/dashboard/monitors")
	}

	// Check if the user owns this monitor
	monitor, err := h.monitorService.GetByIDAndUser(id, user.Id)
	if monitor == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/dashboard/monitors")
	}

	if err := h.monitorService.Delete(id, user.Id); err != nil {
		log2.Errorf("Failed to delete monitor: %v", err)
	}

	return c.Redirect(http.StatusSeeOther, "/dashboard/monitors")
}
