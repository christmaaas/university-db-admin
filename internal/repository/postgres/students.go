package postgres

import (
	"context"
	"log"
	"university-db-admin/internal/domain"
	"university-db-admin/internal/dto"
	"university-db-admin/internal/repository"

	"github.com/jackc/pgx/v5"
)

type studentsRepository struct {
	db *pgx.Conn
}

func NewStudentsRepository(db *pgx.Conn) repository.Students {
	return &studentsRepository{
		db: db,
	}
}

func (s *studentsRepository) Create(ctx context.Context, stud domain.Student) error {
	sql := `
		INSERT INTO public.students (name, passport, employee_id, group_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := s.db.QueryRow(ctx, sql,
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
	err := s.db.QueryRow(ctx, sql, id).Scan(
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

	rows, err := s.db.Query(ctx, sql)
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

	rows, err := s.db.Query(ctx, sql, name)
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
	err := s.db.QueryRow(ctx, sql, passport).Scan(
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

	rows, err := s.db.Query(ctx, sql, id)
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

	rows, err := s.db.Query(ctx, sql, id)
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

func (r *studentsRepository) FindAllWithNoCurator(ctx context.Context) ([]dto.StudentNoCuratorDTO, error) {
	sql := `
        SELECT students.name, 
            students.passport,
            students.group_id
        FROM public.students
        WHERE NOT students.employee_id IS NOT NULL
    `

	log.Println("executing sql:", sql)

	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	var result []dto.StudentNoCuratorDTO
	for rows.Next() {
		var dto dto.StudentNoCuratorDTO
		err := rows.Scan(
			&dto.Name,
			&dto.Passport,
			&dto.GroupID,
		)
		if err != nil {
			return nil, handlePgError(err)
		}
		result = append(result, dto)
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *studentsRepository) FindAllByMiddlename(ctx context.Context, m string) ([]dto.StudentByNameDTO, error) {
	sql := `
		SELECT students.name, students.passport
		FROM public.students
		WHERE students.name LIKE $1
	`

	log.Println("executing sql:", sql)

	pattern := "%" + m
	rows, err := r.db.Query(ctx, sql, pattern)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	var result []dto.StudentByNameDTO
	for rows.Next() {
		var dto dto.StudentByNameDTO
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

func (r *studentsRepository) FindAllGroupCombs(ctx context.Context) ([]dto.StudentGroupCombDTO, error) {
	sql := `
		SELECT students.name, groups.number
		FROM public.students
		CROSS JOIN public.groups
	`

	log.Println("executing sql:", sql)

	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	var result []dto.StudentGroupCombDTO
	for rows.Next() {
		var dto dto.StudentGroupCombDTO
		err := rows.Scan(
			&dto.StudentName,
			&dto.GroupNumber,
		)
		if err != nil {
			return nil, handlePgError(err)
		}
		result = append(result, dto)
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *studentsRepository) FindAllWithCurators(ctx context.Context) ([]dto.StudentCuratorDTO, error) {
	sql := `
		SELECT students.name,
			students.passport,
			employees.name,
			employees.passport
		FROM public.students
		LEFT OUTER JOIN employees ON students.employee_id = employees.id;
	`

	log.Println("executing sql:", sql)

	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	var result []dto.StudentCuratorDTO
	for rows.Next() {
		var dto dto.StudentCuratorDTO
		err := rows.Scan(
			&dto.StudentName,
			&dto.StudentPassport,
			&dto.CuratorName,
			&dto.CuratorPassport,
		)
		if err != nil {
			return nil, handlePgError(err)
		}
		result = append(result, dto)
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *studentsRepository) FindWithAllCurators(ctx context.Context) ([]dto.StudentCuratorDTO, error) {
	sql := `
		SELECT students.name,
			students.passport,
			employees.name,
			employees.passport
		FROM public.students
		RIGHT OUTER JOIN employees ON students.employee_id = employees.id;
	`

	log.Println("executing sql:", sql)

	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	var result []dto.StudentCuratorDTO
	for rows.Next() {
		var dto dto.StudentCuratorDTO
		err := rows.Scan(
			&dto.StudentName,
			&dto.StudentPassport,
			&dto.CuratorName,
			&dto.CuratorPassport,
		)
		if err != nil {
			return nil, handlePgError(err)
		}
		result = append(result, dto)
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *studentsRepository) FindAllPairsWithCurator(ctx context.Context) ([]dto.StudentCuratorDTO, error) {
	sql := `
		SELECT students.name,
			students.passport,
			employees.name,
			employees.passport
		FROM public.students
		FULL OUTER JOIN employees ON students.employee_id = employees.id;
	`

	log.Println("executing sql:", sql)

	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	var result []dto.StudentCuratorDTO
	for rows.Next() {
		var dto dto.StudentCuratorDTO
		err := rows.Scan(
			&dto.StudentName,
			&dto.StudentPassport,
			&dto.CuratorName,
			&dto.CuratorPassport,
		)
		if err != nil {
			return nil, handlePgError(err)
		}
		result = append(result, dto)
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *studentsRepository) FindAllUppercaseWithLength(ctx context.Context) ([]dto.StudentNameStatDTO, error) {
	sql := `
		SELECT students.id,
			UPPER(students.name),
			LENGTH(students.name)
		FROM public.students
	`

	log.Println("executing sql:", sql)

	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	var result []dto.StudentNameStatDTO
	for rows.Next() {
		var dto dto.StudentNameStatDTO
		err := rows.Scan(
			&dto.ID,
			&dto.UppercaseName,
			&dto.NameLength,
		)
		if err != nil {
			return nil, handlePgError(err)
		}
		result = append(result, dto)
	}

	log.Println("sql result:", result)
	return result, nil
}

func (s *studentsRepository) Update(ctx context.Context, id uint64, stud domain.Student) error {
	sql := `
		UPDATE public.students
		SET name = $1, passport = $2, employee_id = $3, group_id = $4
		WHERE id = $5
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := s.db.QueryRow(ctx, sql,
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
	err := s.db.QueryRow(ctx, sql, id).Scan(&id)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", id)
	return nil
}
