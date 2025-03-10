package domain

type Subject struct {
	ID          uint64 `validate:"gte=0"`
	Name        string `validate:"required,min=1"`
	Description string `validate:"required,min=1"`
}
