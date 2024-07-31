package domain

type UserRepository interface {
	FindByEmail(email Email) (User, error)
}
