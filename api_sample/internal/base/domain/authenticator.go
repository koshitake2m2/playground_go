package domain

type Authenticator interface {
	Authenticate(s SessionId) (Email, error)
}
