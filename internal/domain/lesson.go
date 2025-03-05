package domain

type Lesson struct {
	ID           uint64
	GroupID      uint64
	SubjectID    uint64
	LessonTypeID uint64
	Week         uint16
	Weekday      uint16
	Room         uint64
}
