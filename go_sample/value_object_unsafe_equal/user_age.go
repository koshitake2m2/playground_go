package main

type UserAge struct {
	ValueObject[int64]
}

func NewUserAge(age int64) UserAge {
	return UserAge{NewValueObject(age)}
}

func NextAge(age UserAge) UserAge {
	return NewUserAge(age.Value() + 1)
}

type Year struct {
	ValueObject[int64]
}

func NewYear(age int64) Year {
	return Year{NewValueObject(age)}
}
