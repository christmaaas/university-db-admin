package domain

type Group struct {
	ID     uint64 `validate:"gte=0"`
	Number uint64 `validate:"required,gt=0"`
}
