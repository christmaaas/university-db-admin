package domain

type Position struct {
	ID   uint64 `validate:"gte=0"`
	Name string `validate:"required,min=1"`
}
