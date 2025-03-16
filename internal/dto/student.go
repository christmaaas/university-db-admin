package dto

type StudentNoCuratorDTO struct {
	Name     string
	Passport string
	GroupID  uint64
}

type StudentByNameDTO struct {
	Name     string
	Passport string
}

type StudentGroupCombDTO struct {
	StudentName string
	GroupNumber uint64
}

type StudentNameStatDTO struct {
	ID            uint64
	UppercaseName string
	NameLength    uint64
}
