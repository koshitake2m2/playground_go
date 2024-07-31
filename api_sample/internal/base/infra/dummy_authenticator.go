package infra

import (
	"api_sample/internal/base/domain"
)

type DummyAuthenticator struct {
}

// DBアクセスの代わりにダミーな処理をする
func (da DummyAuthenticator) Authenticate(s domain.SessionId) (domain.Email, error) {
	switch s.ToString() {
	case "SESSION_ID_USER_1":
		return domain.NewEmail(dummyUser1Email), nil
	case "SESSION_ID_USER_2":
		return domain.NewEmail(dummyUser2Email), nil
	default:
		return domain.Email{}, domain.AuthenticationError{}
	}
}
