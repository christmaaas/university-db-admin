package domain

type EmployeeSubject struct {
	EmployeeID uint64 `validate:"required,gt=0"`
	SubjectID  uint64 `validate:"required,gt=0"`
}
