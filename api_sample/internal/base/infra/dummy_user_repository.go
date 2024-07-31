package infra

import (
	"api_sample/internal/base/domain"
)

type DummyUserRepository struct {
}

func (d DummyUserRepository) FindByEmail(email domain.Email) (domain.User, error) {
	switch email.ToString() {
	case dummyUser1Email:
		return domain.User{UserId: domain.NewUserId(1), Email: email}, nil
	case dummyUser2Email:
		return domain.User{UserId: domain.NewUserId(2), Email: email}, nil
	default:
		return domain.User{}, domain.UserNotFoundError{}
	}
}
