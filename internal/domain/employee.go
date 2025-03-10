package domain

type Employee struct {
	ID         uint64 `validate:"gte=0"`
	Name       string `validate:"required,min=5"`
	Passport   string `validate:"required,len=9"`
	PositionID uint64 `validate:"required,gt=0"`
}
