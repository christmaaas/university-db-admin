package postgres

import (
	"context"
	"log"
	"university-db-admin/internal/domain"
	"university-db-admin/internal/dto"
	"university-db-admin/internal/repository"

	"github.com/jackc/pgx/v5"
)

type employeesRepository struct {
	db *pgx.Conn
}

func NewEmployeesRepository(db *pgx.Conn) repository.Employees {
	return &employeesRepository{
		db: db,
	}
}

func (e *employeesRepository) Create(ctx context.Context, emp domain.Employee) error {
	sql := `
		INSERT INTO public.employees (name, passport, position_id)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := e.db.QueryRow(ctx, sql,
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
	err := e.db.QueryRow(ctx, sql, id).Scan(
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

	rows, err := e.db.Query(ctx, sql)
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

	rows, err := e.db.Query(ctx, sql, name)
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
	err := e.db.QueryRow(ctx, sql, passport).Scan(
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

	rows, err := e.db.Query(ctx, sql, position)
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

func (r *employeesRepository) FindAllNamePassport(ctx context.Context) ([]dto.EmployeeDTO, error) {
	sql := `
        SELECT employees.name, employees.passport
        FROM public.employees
    `

	log.Println("executing sql:", sql)

	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	var result []dto.EmployeeDTO
	for rows.Next() {
		var dto dto.EmployeeDTO
		err := rows.Scan(
			&dto.Name,
			&dto.Passport,
		)
		if err != nil {
			return nil, handlePgError(err)
		}
		result = append(result, dto)
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *employeesRepository) FindNamePassportByID(ctx context.Context, id uint64) (dto.EmployeeDTO, error) {
	sql := `
        SELECT employees.name, employees.passport
        FROM public.employees
        WHERE employees.id = $1
    `

	log.Println("executing sql:", sql)

	row := r.db.QueryRow(ctx, sql, id)

	var dto dto.EmployeeDTO
	err := row.Scan(
		&dto.Name,
		&dto.Passport,
	)
	if err != nil {
		return dto, handlePgError(err)
	}

	log.Println("sql result:", dto)
	return dto, nil
}

func (r *employeesRepository) FindAllByPositions(ctx context.Context, firstID, secondID uint64) ([]dto.EmployeePositionDTO, error) {
	sql := `
		SELECT employees.name
		FROM public.employees
		WHERE employees.position_id = $1 OR employees.position_id = $2
	`

	log.Println("executing sql:", sql)

	rows, err := r.db.Query(ctx, sql, firstID, secondID)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	var result []dto.EmployeePositionDTO
	for rows.Next() {
		var dto dto.EmployeePositionDTO
		err := rows.Scan(&dto.Name)
		if err != nil {
			return nil, handlePgError(err)
		}
		result = append(result, dto)
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *employeesRepository) IsTeacher(ctx context.Context, id uint64) (dto.EmployeeRoleDTO, error) {
	const teacherName = "Преподаватель"
	sql := `
		SELECT EXISTS (
			SELECT 1 
			FROM employees e
			INNER JOIN positions p ON e.position_id = p.id
			WHERE e.id = $1 AND p.name = $2
		)
	`

	var dto dto.EmployeeRoleDTO
	err := r.db.QueryRow(ctx, sql, id, teacherName).Scan(&dto.IsTeacher)
	if err != nil {
		return dto, handlePgError(err)
	}

	log.Println("sql result:", dto)
	return dto, nil
}

func (e *employeesRepository) Update(ctx context.Context, id uint64, emp domain.Employee) error {
	sql := `
		UPDATE public.employees
		SET name = $1, passport = $2, position_id = $3
		WHERE id = $4
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := e.db.QueryRow(ctx, sql,
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
	err := e.db.QueryRow(ctx, sql, id).Scan(&id)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", id)
	return nil
}
