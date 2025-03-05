package main

type UserAge struct {
	ValueObject[int64, UserAge]
}

func NewUserAge(age int64) UserAge {
	return UserAge{NewValueObject[int64, UserAge](age)}
}

func NextAge(age UserAge) UserAge {
	return NewUserAge(age.Value() + 1)
}

type Year struct {
	ValueObject[int64, Year]
}

func NewYear(age int64) Year {
	return Year{NewValueObject[int64, Year](age)}
}
