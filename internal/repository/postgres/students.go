package postgres

import (
	"context"
	"log"
	"university-db-admin/internal/domain"
	"university-db-admin/internal/repository"

	"github.com/jackc/pgx/v5"
)

type studentsRepository struct {
	dbclient *pgx.Conn
}

func NewStudentsRepository(dbclient *pgx.Conn) repository.Students {
	return &studentsRepository{
		dbclient: dbclient,
	}
}

func (s *studentsRepository) Create(ctx context.Context, stud domain.Student) error {
	sql := `
		INSERT INTO public.students (name, passport, employee_id, group_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := s.dbclient.QueryRow(ctx, sql,
		stud.Name,
		stud.Passport,
		stud.EmployeeID,
		stud.GroupID,
	).Scan(&stud.ID)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", stud.ID)
	return nil
}

func (s *studentsRepository) FindOne(ctx context.Context, id uint64) (domain.Student, error) {
	sql := `
		SELECT id, name, passport, employee_id, group_id
		FROM public.students
		WHERE id = $1
	`

	var stud domain.Student
	log.Println("executing sql:", sql)
	err := s.dbclient.QueryRow(ctx, sql, id).Scan(
		&stud.ID,
		&stud.Name,
		&stud.Passport,
		&stud.EmployeeID,
		&stud.GroupID,
	)
	if err != nil {
		return domain.Student{}, handlePgError(err)
	}

	log.Println("sql result:", stud)
	return stud, nil
}

func (s *studentsRepository) FindAll(ctx context.Context) ([]domain.Student, error) {
	sql := `
		SELECT id, name, passport, employee_id, group_id
		FROM public.students
	`

	var students []domain.Student
	log.Println("executing sql:", sql)

	rows, err := s.dbclient.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var stud domain.Student
		err := rows.Scan(
			&stud.ID,
			&stud.Name,
			&stud.Passport,
			&stud.EmployeeID,
			&stud.GroupID,
		)
		if err != nil {
			return nil, handlePgError(err)
		}
		students = append(students, stud)
	}

	log.Println("sql result:", students)
	return students, nil
}

func (s *studentsRepository) FindByName(ctx context.Context, name string) ([]domain.Student, error) {
	sql := `
		SELECT id, name, passport, employee_id, group_id
		FROM public.students
		WHERE name = $1
	`

	var students []domain.Student
	log.Println("executing sql:", sql)

	rows, err := s.dbclient.Query(ctx, sql, name)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var stud domain.Student
		err := rows.Scan(
			&stud.ID,
			&stud.Name,
			&stud.Passport,
			&stud.EmployeeID,
			&stud.GroupID,
		)
		if err != nil {
			return nil, handlePgError(err)
		}
		students = append(students, stud)
	}

	log.Println("sql result:", students)
	return students, nil
}

func (s *studentsRepository) FindByPassport(ctx context.Context, passport string) (domain.Student, error) {
	sql := `
		SELECT id, name, passport, employee_id, group_id
		FROM public.students
		WHERE passport = $1
	`

	var stud domain.Student
	log.Println("executing sql:", sql)
	err := s.dbclient.QueryRow(ctx, sql, passport).Scan(
		&stud.ID,
		&stud.Name,
		&stud.Passport,
		&stud.EmployeeID,
		&stud.GroupID,
	)
	if err != nil {
		return domain.Student{}, handlePgError(err)
	}

	log.Println("sql result:", stud)
	return stud, nil
}

func (s *studentsRepository) FindByEmployeeID(ctx context.Context, id uint64) ([]domain.Student, error) {
	sql := `
		SELECT id, name, passport, employee_id, group_id
		FROM public.students
		WHERE employee_id = $1
	`

	var students []domain.Student
	log.Println("executing sql:", sql)

	rows, err := s.dbclient.Query(ctx, sql, id)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var stud domain.Student
		err := rows.Scan(
			&stud.ID,
			&stud.Name,
			&stud.Passport,
			&stud.EmployeeID,
			&stud.GroupID,
		)
		if err != nil {
			return nil, handlePgError(err)
		}
		students = append(students, stud)
	}

	log.Println("sql result:", students)
	return students, nil
}

func (s *studentsRepository) FindByGroupID(ctx context.Context, id uint64) ([]domain.Student, error) {
	sql := `
		SELECT id, name, passport, employee_id, group_id
		FROM public.students
		WHERE group_id = $1
	`

	var students []domain.Student
	log.Println("executing sql:", sql)

	rows, err := s.dbclient.Query(ctx, sql, id)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var stud domain.Student
		err := rows.Scan(
			&stud.ID,
			&stud.Name,
			&stud.Passport,
			&stud.EmployeeID,
			&stud.GroupID,
		)
		if err != nil {
			return nil, handlePgError(err)
		}
		students = append(students, stud)
	}

	log.Println("sql result:", students)
	return students, nil
}

func (s *studentsRepository) Update(ctx context.Context, id uint64, stud domain.Student) error {
	sql := `
		UPDATE public.students
		SET name = $1, passport = $2, employee_id = $3, group_id = $4
		WHERE id = $5
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := s.dbclient.QueryRow(ctx, sql,
		stud.Name,
		stud.Passport,
		stud.EmployeeID,
		stud.GroupID,
		id,
	).Scan(&id)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", id)
	return nil
}

func (s *studentsRepository) Delete(ctx context.Context, id uint64) error {
	sql := `
		DELETE FROM public.students
		WHERE id = $1
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := s.dbclient.QueryRow(ctx, sql, id).Scan(&id)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", id)
	return nil
}
