package dto

import "time"

type MarkBySubjectDTO struct {
	StudentID uint64
	Mark      uint16
	Date      time.Time
}

type SortedMarkDTO struct {
	StudentID uint64
	Mark      uint16
	Date      time.Time
}
