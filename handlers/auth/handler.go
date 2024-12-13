package auth

import (
	"fmt"
	"github.com/200sh/200sh-dashboard/middleware"
	"github.com/200sh/200sh-dashboard/models"
	"github.com/200sh/200sh-dashboard/views/auth"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type Handler struct {
	HankoApiUrl string
	UserService models.UserService
}

func (ah *Handler) LoginPageHandler(c echo.Context) error {
	isLoggedIn := c.Get(middleware.IsLoggedInKey).(bool)
	if isLoggedIn {
		return c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
	}
	return auth.Login(ah.HankoApiUrl).Render(c.Response().Writer)
}

func (h *Handler) UserSetupPage(c echo.Context) error {
	if !c.Get(middleware.IsLoggedInKey).(bool) {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	return auth.UserSetup().Render(c.Response().Writer)
}

type UserForm struct {
	FirstName string `form:"first-name"`
	LastName  string `form:"last-name"`
}

func (h *Handler) UserSetupForm(c echo.Context) error {
	// Get the form data
	var uf UserForm
	err := c.Bind(&uf)
	if err != nil {
		return c.String(http.StatusBadRequest, "Not valid form")
	}

	// Validate the data
	// TODO: need to return something that can be used by the form? A new page with some error message

	// Create new User and insert into db
	// TODO: Get user Data from Hanko api
	// TODO: Need to get the token from the cookie

	// User Hanko data for
	user := models.User{
		ProviderId: "", // From Hanko cookie and/or user data
		Provider:   "hanko",
		Email:      "", // From Hanko User data
		Name:       fmt.Sprintf("%s %s", uf.FirstName, uf.LastName),
		Status:     models.UserStatusActive,
		CreatedAt:  time.Time{}, // From Hanko User data
	}
	err = h.UserService.CreateUser(&user)
	if err != nil {
		return err
	}

	// Redirect to dashboard
	return c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
}
