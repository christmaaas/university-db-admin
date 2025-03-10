package domain

type Lesson struct {
	ID           uint64 `validate:"gte=0"`
	GroupID      uint64 `validate:"required,gt=0"`
	SubjectID    uint64 `validate:"required,gt=0"`
	LessonTypeID uint64 `validate:"required,gt=0"`
	Week         uint16 `validate:"required,gt=0"`
	Weekday      uint16 `validate:"required,gt=0"`
	Room         uint64 `validate:"required,gt=0"`
}
