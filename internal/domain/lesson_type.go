package domain

type LessonType struct {
	ID   uint64 `validate:"gte=0"`
	Name string `validate:"required,len=2"`
}
