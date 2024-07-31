package presentation

import (
	"api_sample/internal/base/domain"

	"github.com/labstack/echo/v4"
)

type AuthenticationHelper struct {
	Authenticator  domain.Authenticator
	UserRepository domain.UserRepository
}

const COOKIE_NAME_SESSION_ID = "SESSION_ID"

func (ah AuthenticationHelper) Authenticate(c echo.Context) (domain.User, error) {
	cookie, err := c.Cookie(COOKIE_NAME_SESSION_ID)
	if err != nil {
		return domain.User{}, err
	}
	sessionId := domain.NewSessionId(cookie.Value)
	email, err := ah.Authenticator.Authenticate(sessionId)
	if err != nil {
		return domain.User{}, err
	}

	users, err := ah.UserRepository.FindByEmail(email)
	if err != nil {
		return domain.User{}, err
	}

	return users, nil
}
