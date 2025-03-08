package postgres

import (
	"context"
	"log"
	"university-db-admin/internal/domain"
	"university-db-admin/internal/repository"

	"github.com/jackc/pgx/v5"
)

type employeesRepository struct {
	dbclient *pgx.Conn
}

func NewEmployeesRepository(dbclient *pgx.Conn) repository.Employees {
	return &employeesRepository{
		dbclient: dbclient,
	}
}

func (e *employeesRepository) Create(ctx context.Context, emp domain.Employee) error {
	sql := `
		INSERT INTO public.employees (name, passport, position_id)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := e.dbclient.QueryRow(ctx, sql,
		emp.Name,
		emp.Passport,
		emp.PositionID,
	).Scan(&emp.ID)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", emp.ID)
	return nil
}

func (e *employeesRepository) FindOne(ctx context.Context, id uint64) (domain.Employee, error) {
	sql := `
		SELECT id, name, passport, position_id 
		FROM public.employees
		WHERE id = $1
	`

	var emp domain.Employee
	log.Println("executing sql:", sql)
	err := e.dbclient.QueryRow(ctx, sql, id).Scan(
		&emp.ID,
		&emp.Name,
		&emp.Passport,
		&emp.PositionID,
	)
	if err != nil {
		return domain.Employee{}, handlePgError(err)
	}

	log.Println("sql result:", emp)
	return emp, nil
}

func (e *employeesRepository) FindAll(ctx context.Context) ([]domain.Employee, error) {
	sql := `
		SELECT id, name, passport, position_id 
		FROM public.employees
	`

	var emps []domain.Employee
	log.Println("executing sql:", sql)

	rows, err := e.dbclient.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var emp domain.Employee
		err := rows.Scan(
			&emp.ID,
			&emp.Name,
			&emp.Passport,
			&emp.PositionID,
		)
		if err != nil {
			return nil, handlePgError(err)
		}
		emps = append(emps, emp)
	}

	log.Println("sql result:", emps)
	return emps, nil
}

func (e *employeesRepository) FindByName(ctx context.Context, name string) ([]domain.Employee, error) {
	sql := `
		SELECT id, name, passport, position_id 
		FROM public.employees
		WHERE name = $1
	`

	var emps []domain.Employee
	log.Println("executing sql:", sql)

	rows, err := e.dbclient.Query(ctx, sql, name)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var emp domain.Employee
		err := rows.Scan(
			&emp.ID,
			&emp.Name,
			&emp.Passport,
			&emp.PositionID,
		)
		if err != nil {
			return nil, handlePgError(err)
		}
		emps = append(emps, emp)
	}

	log.Println("sql result:", emps)
	return emps, nil
}

func (e *employeesRepository) FindByPassport(ctx context.Context, passport string) (domain.Employee, error) {
	sql := `
		SELECT id, name, passport, position_id 
		FROM public.employees
		WHERE passport = $1
	`

	var emp domain.Employee
	log.Println("executing sql:", sql)
	err := e.dbclient.QueryRow(ctx, sql, passport).Scan(
		&emp.ID,
		&emp.Name,
		&emp.Passport,
		&emp.PositionID,
	)
	if err != nil {
		return domain.Employee{}, handlePgError(err)
	}

	log.Println("sql result:", emp)
	return emp, nil
}

func (e *employeesRepository) FindByPosition(ctx context.Context, position uint64) ([]domain.Employee, error) {
	sql := `
		SELECT id, name, passport, position_id 
		FROM public.employees
		WHERE position_id = $1
	`

	var emps []domain.Employee
	log.Println("executing sql:", sql)

	rows, err := e.dbclient.Query(ctx, sql, position)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var emp domain.Employee
		err := rows.Scan(
			&emp.ID,
			&emp.Name,
			&emp.Passport,
			&emp.PositionID,
		)
		if err != nil {
			return nil, handlePgError(err)
		}
		emps = append(emps, emp)
	}

	log.Println("sql result:", emps)
	return emps, nil
}

func (e *employeesRepository) Update(ctx context.Context, id uint64, emp domain.Employee) error {
	sql := `
		UPDATE public.employees
		SET name = $1, passport = $2, position_id = $3
		WHERE id = $4
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := e.dbclient.QueryRow(ctx, sql,
		emp.Name,
		emp.Passport,
		emp.PositionID,
		id,
	).Scan(&id)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", id)
	return nil
}

func (e *employeesRepository) Delete(ctx context.Context, id uint64) error {
	sql := `
		DELETE FROM public.employees
		WHERE id = $1
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := e.dbclient.QueryRow(ctx, sql, id).Scan(&id)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", id)
	return nil
}
