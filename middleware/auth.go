package middleware

import (
	"database/sql"
	"errors"
	"github.com/200sh/200sh-dashboard/hanko"
	"github.com/200sh/200sh-dashboard/models/services"
	"github.com/labstack/echo/v4"
	log2 "github.com/labstack/gommon/log"
	"net/http"
)

const (
	IsLoggedInKey = "IsLoggedIn"
	UserIDKey     = "User"
)

type AuthMiddleware struct {
	hanko       *hanko.Hanko
	userService services.UserService
}

func NewAuthMiddleware(hankoClient *hanko.Hanko, userService services.UserService) *AuthMiddleware {
	return &AuthMiddleware{
		hanko:       hankoClient,
		userService: userService,
	}
}

func (am AuthMiddleware) IsLoggedInEnriched() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			_, err := am.hanko.ValidateHankoCookie(c)
			c.Set(IsLoggedInKey, err == nil)
			return next(c)
		}
	}
}

// AuthRequired
// Check the auth cookie 'hanko' and checks if it is valid
// Then fetches the user data from the db if it exists and attaches it to the context
// If there is no user in the db we redirect to the user-setup-flow
func (am AuthMiddleware) AuthRequired() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := am.hanko.ValidateHankoCookie(c)
			if errors.Is(err, http.ErrNoCookie) {
				log2.Info("No 'hanko' cookie set")
				return c.Redirect(http.StatusTemporaryRedirect, "/login")
			}
			if err != nil {
				log2.Info("Failed to get valid token from 'hanko' cookie")
				return c.Redirect(http.StatusTemporaryRedirect, "/login")
			}

			// Try fetch the User information from the db
			user, err := am.userService.GetBySubjectID(token.Subject())
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					// Redirect to user setup flow
					return c.Redirect(http.StatusTemporaryRedirect, "/user/setup")
				}

				log2.Fatalf("Error fetching user from db: %s", err)
			}

			c.Set(UserIDKey, user)

			return next(c)
		}
	}
}
