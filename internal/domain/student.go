package domain

type Student struct {
	ID         uint64 `validate:"gte=0"`
	Name       string `validate:"required,min=5"`
	Passport   string `validate:"required,len=9"`
	EmployeeID uint64 `validate:"required,gt=0"`
	GroupID    uint64 `validate:"required,gt=0"`
}
