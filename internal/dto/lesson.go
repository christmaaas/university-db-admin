package dto

type LessonScheduleDTO struct {
	GroupNumber uint64
	Subject     string
	LessonType  string
	Room        uint64
	Week        uint16
	Weekday     uint16
}
