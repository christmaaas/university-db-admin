package repository

import (
	"context"
	"university-db-admin/internal/domain"
	"university-db-admin/internal/dto"
)

type Repository struct {
	Employees         Employees
	Groups            Groups
	LessonTypes       LessonTypes
	Lessons           Lessons
	Marks             Marks
	Positions         Positions
	Students          Students
	Subjects          Subjects
	EmployeesSubjects EmployeesSubjects
}

type Employees interface {
	Create(ctx context.Context, emp domain.Employee) error
	FindOne(ctx context.Context, id uint64) (domain.Employee, error)
	FindAll(ctx context.Context) ([]domain.Employee, error)
	FindByName(ctx context.Context, name string) ([]domain.Employee, error)
	FindByPassport(ctx context.Context, passport string) (domain.Employee, error)
	FindByPosition(ctx context.Context, position uint64) ([]domain.Employee, error)
	FindAllNamePassport(ctx context.Context) ([]dto.EmployeeDTO, error)
	FindNamePassportByID(ctx context.Context, id uint64) (dto.EmployeeDTO, error)
	FindAllByPositions(ctx context.Context, firstID, secondID uint64) ([]dto.EmployeePositionDTO, error)
	IsTeacher(ctx context.Context, id uint64) (dto.EmployeeRoleDTO, error)
	Update(ctx context.Context, id uint64, emp domain.Employee) error
	Delete(ctx context.Context, id uint64) error
}

type Groups interface {
	Create(ctx context.Context, grp domain.Group) error
	FindOne(ctx context.Context, id uint64) (domain.Group, error)
	FindAll(ctx context.Context) ([]domain.Group, error)
	FindByNumber(ctx context.Context, num uint64) (domain.Group, error)
	Update(ctx context.Context, id uint64, grp domain.Group) error
	Delete(ctx context.Context, id uint64) error
}

type LessonTypes interface {
	Create(ctx context.Context, lsn domain.LessonType) error
	FindOne(ctx context.Context, id uint64) (domain.LessonType, error)
	FindAll(ctx context.Context) ([]domain.LessonType, error)
	FindByName(ctx context.Context, name string) (domain.LessonType, error)
	Update(ctx context.Context, id uint64, lsn domain.LessonType) error
	Delete(ctx context.Context, id uint64) error
}

type Lessons interface {
	Create(ctx context.Context, lsn domain.Lesson) error
	FindOne(ctx context.Context, id uint64) (domain.Lesson, error)
	FindAll(ctx context.Context) ([]domain.Lesson, error)
	FindByGroupID(ctx context.Context, id uint64) ([]domain.Lesson, error)
	FindBySubjectID(ctx context.Context, id uint64) ([]domain.Lesson, error)
	FindByLessonTypeID(ctx context.Context, id uint64) ([]domain.Lesson, error)
	FindByWeek(ctx context.Context, week uint16) ([]domain.Lesson, error)
	FindByWeekday(ctx context.Context, weekday uint16) ([]domain.Lesson, error)
	FindByRoom(ctx context.Context, room uint64) ([]domain.Lesson, error)
	FindSchedule(ctx context.Context) ([]dto.LessonScheduleDTO, error)
	Update(ctx context.Context, id uint64, emp domain.Lesson) error
	Delete(ctx context.Context, id uint64) error
}

type Marks interface {
	Create(ctx context.Context, mark domain.Mark) error
	FindOne(ctx context.Context, id uint64) (domain.Mark, error)
	FindAll(ctx context.Context) ([]domain.Mark, error)
	FindByEmployeeID(ctx context.Context, id uint64) ([]domain.Mark, error)
	FindByStudentID(ctx context.Context, id uint64) ([]domain.Mark, error)
	FindBySubjectID(ctx context.Context, id uint64) ([]domain.Mark, error)
	FindByMark(ctx context.Context, mark uint16) ([]domain.Mark, error)
	FindByDate(ctx context.Context, date string) ([]domain.Mark, error)
	FindAllBySubject(ctx context.Context, id uint64, m uint16) ([]dto.MarkBySubjectDTO, error)
	FindAllSorted(ctx context.Context) ([]dto.SortedMarkDTO, error)
	Update(ctx context.Context, id uint64, mark domain.Mark) error
	Delete(ctx context.Context, id uint64) error
}

type Positions interface {
	Create(ctx context.Context, pos domain.Position) error
	FindOne(ctx context.Context, id uint64) (domain.Position, error)
	FindAll(ctx context.Context) ([]domain.Position, error)
	FindByName(ctx context.Context, name string) (domain.Position, error)
	Update(ctx context.Context, id uint64, pos domain.Position) error
	Delete(ctx context.Context, id uint64) error
}

type Students interface {
	Create(ctx context.Context, stud domain.Student) error
	FindOne(ctx context.Context, id uint64) (domain.Student, error)
	FindAll(ctx context.Context) ([]domain.Student, error)
	FindByName(ctx context.Context, name string) ([]domain.Student, error)
	FindByPassport(ctx context.Context, passport string) (domain.Student, error)
	FindByEmployeeID(ctx context.Context, id uint64) ([]domain.Student, error)
	FindByGroupID(ctx context.Context, id uint64) ([]domain.Student, error)
	FindAllWithNoCurator(ctx context.Context) ([]dto.StudentNoCuratorDTO, error)
	FindAllByMiddlename(ctx context.Context, m string) ([]dto.StudentByNameDTO, error)
	FindAllGroupCombs(ctx context.Context) ([]dto.StudentGroupCombDTO, error)
	FindAllWithCurators(ctx context.Context) ([]dto.StudentCuratorDTO, error)
	FindWithAllCurators(ctx context.Context) ([]dto.StudentCuratorDTO, error)
	FindAllPairsWithCurator(ctx context.Context) ([]dto.StudentCuratorDTO, error)
	FindAllUppercaseWithLength(ctx context.Context) ([]dto.StudentNameStatDTO, error)
	Update(ctx context.Context, id uint64, stud domain.Student) error
	Delete(ctx context.Context, id uint64) error
}

type Subjects interface {
	Create(ctx context.Context, sbj domain.Subject) error
	FindOne(ctx context.Context, id uint64) (domain.Subject, error)
	FindAll(ctx context.Context) ([]domain.Subject, error)
	FindByName(ctx context.Context, name string) (domain.Subject, error)
	FindAllSorted(ctx context.Context) ([]dto.SortedSubjectDTO, error)
	Update(ctx context.Context, id uint64, sbj domain.Subject) error
	Delete(ctx context.Context, id uint64) error
}

type EmployeesSubjects interface {
	Create(ctx context.Context, es domain.EmployeeSubject) error
	FindAll(ctx context.Context) ([]domain.EmployeeSubject, error)
	FindByEmployeeID(ctx context.Context, id uint64) ([]domain.EmployeeSubject, error)
	FindBySubjectID(ctx context.Context, id uint64) ([]domain.EmployeeSubject, error)
	Update(ctx context.Context, eid uint64, sid uint64, es domain.EmployeeSubject) error
	Delete(ctx context.Context, eid uint64, sid uint64) error
}
