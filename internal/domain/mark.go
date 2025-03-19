package domain

import "time"

type Mark struct {
	ID         uint64    `validate:"gte=0"`
	EmployeeID uint64    `validate:"required,gt=0"`
	StudentID  uint64    `validate:"required,gt=0"`
	SubjectID  uint64    `validate:"required,gt=0"`
	Mark       uint16    `validate:"required,gt=0"`
	Date       time.Time `validate:"required"`
}
