package domain

type UserId struct {
	value int
}

func (v UserId) ToInt() int {
	return v.value
}

func NewUserId(id int) UserId {
	return UserId{value: id}
}

type Email struct {
	value string
}

func (v Email) ToString() string {
	return v.value
}

func NewEmail(email string) Email {
	return Email{value: email}
}

type User struct {
	UserId UserId
	Email  Email
}

type UserNotFoundError struct {
}

func (v UserNotFoundError) Error() string {
	return "user not found"
}
