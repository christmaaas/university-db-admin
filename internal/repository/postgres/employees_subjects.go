package postgres

import (
	"context"
	"log"
	"university-db-admin/internal/domain"
	"university-db-admin/internal/repository"

	"github.com/jackc/pgx/v5"
)

type employeesSubjectsRepository struct {
	db *pgx.Conn
}

func NewEmployeesSubjectsRepository(db *pgx.Conn) repository.EmployeesSubjects {
	return &employeesSubjectsRepository{
		db: db,
	}
}

func (s *employeesSubjectsRepository) Create(ctx context.Context, es domain.EmployeeSubject) error {
	sql := `
		INSERT INTO public.employees_subjects (employee_id, subject_id)
		VALUES ($1, $2)
		RETURNING employee_id, subject_id
	`

	log.Println("executing sql:", sql)
	err := s.db.QueryRow(ctx, sql, es.EmployeeID, es.SubjectID).Scan(
		&es.EmployeeID,
		&es.SubjectID,
	)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", es.EmployeeID, es.SubjectID)
	return nil
}

func (s *employeesSubjectsRepository) FindAll(ctx context.Context) ([]domain.EmployeeSubject, error) {
	sql := `
		SELECT employee_id, subject_id
		FROM public.employees_subjects
	`

	var empSbjs []domain.EmployeeSubject
	log.Println("executing sql:", sql)

	rows, err := s.db.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var es domain.EmployeeSubject
		err := rows.Scan(&es.EmployeeID, &es.SubjectID)
		if err != nil {
			return nil, handlePgError(err)
		}
		empSbjs = append(empSbjs, es)
	}

	log.Println("sql result:", empSbjs)
	return empSbjs, nil
}

func (s *employeesSubjectsRepository) FindByEmployeeID(ctx context.Context, id uint64) ([]domain.EmployeeSubject, error) {
	sql := `
		SELECT employee_id, subject_id
		FROM public.employees_subjects
		WHERE employee_id = $1
	`

	var empSbjs []domain.EmployeeSubject
	log.Println("executing sql:", sql)

	rows, err := s.db.Query(ctx, sql, id)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var es domain.EmployeeSubject
		err := rows.Scan(&es.EmployeeID, &es.SubjectID)
		if err != nil {
			return nil, handlePgError(err)
		}
		empSbjs = append(empSbjs, es)
	}

	log.Println("sql result:", empSbjs)
	return empSbjs, nil
}

func (s *employeesSubjectsRepository) FindBySubjectID(ctx context.Context, id uint64) ([]domain.EmployeeSubject, error) {
	sql := `
		SELECT employee_id, subject_id
		FROM public.employees_subjects
		WHERE subject_id = $1
	`

	var empSbjs []domain.EmployeeSubject
	log.Println("executing sql:", sql)

	rows, err := s.db.Query(ctx, sql, id)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var es domain.EmployeeSubject
		err := rows.Scan(&es.EmployeeID, &es.SubjectID)
		if err != nil {
			return nil, handlePgError(err)
		}
		empSbjs = append(empSbjs, es)
	}

	log.Println("sql result:", empSbjs)
	return empSbjs, nil
}

func (s *employeesSubjectsRepository) Update(ctx context.Context, eid uint64, sid uint64, es domain.EmployeeSubject) error {
	sql := `
		UPDATE public.employees_subjects
		SET employee_id = $1, subject_id = $2
		WHERE employee_id = $3 AND subject_id = $4
		RETURNING employee_id, subject_id
	`

	log.Println("executing sql:", sql)
	err := s.db.QueryRow(ctx, sql,
		es.EmployeeID,
		es.SubjectID,
		eid,
		sid,
	).Scan(&eid, &sid)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", eid, sid)
	return nil
}

func (s *employeesSubjectsRepository) Delete(ctx context.Context, eid uint64, sid uint64) error {
	sql := `
		DELETE FROM public.employees_subjects
		WHERE employee_id = $1 AND subject_id = $2
		RETURNING employee_id, subject_id
	`

	log.Println("executing sql:", sql)
	err := s.db.QueryRow(ctx, sql, eid, sid).Scan(&eid, &sid)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", eid, sid)
	return nil
}
