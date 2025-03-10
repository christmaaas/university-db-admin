package domain

type Mark struct {
	ID         uint64 `validate:"gte=0"`
	EmployeeID uint64 `validate:"required,gt=0"`
	StudentID  uint64 `validate:"required,gt=0"`
	SubjectID  uint64 `validate:"required,gt=0"`
	Mark       uint16 `validate:"required,gt=0"`
	Date       string `validate:"required,datetime=2006-01-02"`
}
