package hanko

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

type Hanko struct {
	HankoApiUrl string
}

type Email struct {
	Address    string
	IsPrimary  bool
	IsVerified bool
}

func NewEmail(emailMap map[string]interface{}) (Email, error) {
	var (
		address    string
		isPrimary  bool
		isVerified bool
	)

	address, ok := emailMap["address"].(string)
	if !ok {
		return Email{}, fmt.Errorf("address not found in emailMap")
	}

	isPrimary, ok = emailMap["is_primary"].(bool)
	if !ok {
		return Email{}, fmt.Errorf("is_primary not found in emailMap")
	}

	isVerified, ok = emailMap["is_verified"].(bool)
	if !ok {
		return Email{}, fmt.Errorf("is_verified not found in emailMap")
	}

	return Email{
		Address:    address,
		IsPrimary:  isPrimary,
		IsVerified: isVerified,
	}, nil
}

func New(hankoApiUrl string) Hanko {
	return Hanko{HankoApiUrl: hankoApiUrl}
}

func (h *Hanko) ValidateHankoCookie(c echo.Context) (jwt.Token, error) {
	cookie, err := c.Cookie("hanko") // TODO: CONFIG: Add this to a config
	if err != nil {
		return nil, err
	}

	set, err := jwk.Fetch(
		context.Background(),
		fmt.Sprintf("%v/.well-known/jwks.json", h.HankoApiUrl),
	)
	if err != nil {
		return nil, err
	}

	return jwt.Parse([]byte(cookie.Value), jwt.WithKeySet(set))
}
