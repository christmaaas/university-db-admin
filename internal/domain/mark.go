package domain

type Mark struct {
	ID         uint64
	EmployeeID uint64
	StudentID  uint64
	SubjectID  uint64
	Mark       uint16
	Date       string
}
