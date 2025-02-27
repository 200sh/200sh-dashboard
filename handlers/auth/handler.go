package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/200sh/200sh-dashboard/hanko"
	"github.com/200sh/200sh-dashboard/middleware"
	"github.com/200sh/200sh-dashboard/models"
	"github.com/200sh/200sh-dashboard/models/services"
	"github.com/200sh/200sh-dashboard/views/auth"
	"github.com/labstack/echo/v4"
	log2 "github.com/labstack/gommon/log"
	"net/http"
)

type Handler struct {
	hanko          *hanko.Hanko
	userService    services.UserService
	monitorService services.MonitorService
}

func NewHandler(hankoClient *hanko.Hanko, userService services.UserService, monitorService services.MonitorService) *Handler {
	return &Handler{
		hanko:          hankoClient,
		userService:    userService,
		monitorService: monitorService,
	}
}

func (h *Handler) LoginPageHandler(c echo.Context) error {
	isLoggedIn := c.Get(middleware.IsLoggedInKey).(bool)
	if isLoggedIn {
		return c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
	}
	return auth.Login(h.hanko.HankoApiUrl).Render(c.Response().Writer)
}

func (h *Handler) UserSetupPage(c echo.Context) error {
	if !c.Get(middleware.IsLoggedInKey).(bool) {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	token, err := h.hanko.ValidateHankoCookie(c)
	if err != nil {
		log2.Warnf("Not able to validate 'hanko' cookie, %s", err)
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}

	_, err = h.userService.GetBySubjectID(token.Subject())
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
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

	log2.Infof("got user form data: %s", uf)
	// Validate the data
	// TODO: need to return something that can be used by the form? A new page with some error message

	// Create new User and insert into db
	token, err := h.hanko.ValidateHankoCookie(c)
	if err != nil {
		log2.Warnf("Not able to validate 'hanko' cookie, %s", err)
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	log2.Infof("Got token: %s", token)

	emailData, ok := token.Get("email")
	if !ok {
		log2.Errorf("Could not find email in 'hanko' jwt")
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	log2.Infof("Got email data: %s", emailData)

	ed, ok := emailData.(map[string]interface{})
	if !ok {
		log2.Errorf("Could not convert email data to map")
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	email, err := hanko.NewEmail(ed)
	if err != nil {
		log2.Errorf("Could not create email from data")
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	if !email.IsVerified {
		log2.Info("Email is not verified")
		// Should redirect to login with error message that email is not verified
		return c.Redirect(http.StatusSeeOther, "/login?error=email_not_verified")
	}

	//user, err := h.Repo.CreateUser(h.Ctx, repository.CreateUserParams{
	//	ProviderID: token.Subject(),
	//	Provider:   "hanko",
	//	Name:       fmt.Sprintf("%s %s", uf.FirstName, uf.LastName),
	//	Email:      email.Address,
	//	Status:     models.UserStatusActive,
	//})
	user := models.User{
		ProviderId: token.Subject(),
		Provider:   "hanko",
		Email:      email.Address,
		Name:       fmt.Sprintf("%s %s", uf.FirstName, uf.LastName),
		Status:     models.UserStatusActive,
	}
	err = h.userService.Create(&user)
	if err != nil {
		return err
	}

	// Redirect to dashboard
	return c.Redirect(http.StatusSeeOther, "/dashboard")
}
