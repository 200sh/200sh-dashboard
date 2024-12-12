package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	log2 "github.com/labstack/gommon/log"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
	"net/http"
)

const (
	IsLoggedInKey = "IsLoggedIn"
	UserIDKey     = "User"
)

type AuthMiddleware struct {
	HankoApiUrl string
}

func (am AuthMiddleware) IsLoggedInEnriched() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			_, err := am.validateHankoCookie(c)
			c.Set(IsLoggedInKey, err == nil)
			return next(c)
		}
	}
}

func (am AuthMiddleware) validateHankoCookie(c echo.Context) (jwt.Token, error) {
	cookie, err := c.Cookie("hanko") // TODO: CONFIG: Add this to a config
	if err != nil {
		return nil, err
	}

	set, err := jwk.Fetch(
		context.Background(),
		fmt.Sprintf("%v/.well-known/jwks.json", am.HankoApiUrl),
	)
	if err != nil {
		return nil, err
	}

	return jwt.Parse([]byte(cookie.Value), jwt.WithKeySet(set))
}

func (am AuthMiddleware) AuthRequired() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token, err := am.validateHankoCookie(c)
			if errors.Is(err, http.ErrNoCookie) {
				log2.Info("No 'hanko' cookie set")
				return c.Redirect(http.StatusTemporaryRedirect, "/login")
			}
			if err != nil {
				log2.Info("Failed to get valid token from 'hanko' cookie")
				return c.Redirect(http.StatusTemporaryRedirect, "/login")
			}

			c.Set(UserIDKey, token.Subject())
			return next(c)
		}
	}
}
