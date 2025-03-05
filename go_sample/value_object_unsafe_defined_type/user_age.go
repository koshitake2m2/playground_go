package main

type UserAge ValueObject[int64, UserAge]

func NewUserAge(age int64) UserAge {
	return NewValueObject[int64, UserAge](age)
}

func NextAge(age UserAge) UserAge {
	return NewUserAge(age.Value() + 1)
}

type Year ValueObject[int64, Year]

func NewYear(age int64) Year {
	return NewValueObject[int64, Year](age)
}
