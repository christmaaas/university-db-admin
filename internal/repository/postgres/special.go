package postgres

import (
	"context"
	"log"
	"university-db-admin/internal/dto"
	"university-db-admin/internal/repository"

	"github.com/jackc/pgx/v5"
)

type specialRequestsRepository struct {
	db *pgx.Conn
}

func NewSpecialRequestsRepository(db *pgx.Conn) repository.Special {
	return &specialRequestsRepository{
		db: db,
	}
}

func (r *specialRequestsRepository) GetAllEmployees(ctx context.Context) ([]dto.EmployeeDTO, error) {
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

func (r *specialRequestsRepository) GetEmployeeByID(ctx context.Context, id uint64) (dto.EmployeeDTO, error) {
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

func (r *specialRequestsRepository) GetStudentsNoCurator(ctx context.Context) ([]dto.StudentNoCuratorDTO, error) {
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

func (r *specialRequestsRepository) GetLessonsSchedule(ctx context.Context) ([]dto.LessonScheduleDTO, error) {
	sql := `
		SELECT groups.number,
			subjects.name,
			lesson_types.name,
			lessons.room,
			lessons.week,
			lessons.weekday
		FROM public.lessons
		INNER JOIN public.groups ON lessons.group_id = groups.id
		INNER JOIN public.subjects ON lessons.subject_id = subjects.id
		INNER JOIN public.lesson_types ON lessons.lesson_type_id = lesson_types.id
	`

	log.Println("executing sql:", sql)

	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	var result []dto.LessonScheduleDTO
	for rows.Next() {
		var dto dto.LessonScheduleDTO
		err := rows.Scan(
			&dto.GroupNumber,
			&dto.Subject,
			&dto.LessonType,
			&dto.Room,
			&dto.Week,
			&dto.Weekday,
		)
		if err != nil {
			return nil, handlePgError(err)
		}
		result = append(result, dto)
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *specialRequestsRepository) GetEmployeesByPositions(ctx context.Context, firstID, secondID uint64) ([]dto.EmployeePositionDTO, error) {
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

func (r *specialRequestsRepository) GetMarksBySubject(ctx context.Context, id uint64, m uint16) ([]dto.MarkBySubjectDTO, error) {
	sql := `
		SELECT marks.student_id, marks.mark, marks.date
		FROM public.marks
		WHERE subject_id = $1 AND mark > $2
	`

	log.Println("executing sql:", sql)

	rows, err := r.db.Query(ctx, sql, id, m)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	var result []dto.MarkBySubjectDTO
	for rows.Next() {
		var dto dto.MarkBySubjectDTO
		err := rows.Scan(
			&dto.StudentID,
			&dto.Mark,
			&dto.Date,
		)
		if err != nil {
			return nil, handlePgError(err)
		}
		result = append(result, dto)
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *specialRequestsRepository) GetStudentsByMiddlename(ctx context.Context, m string) ([]dto.StudentByNameDTO, error) {
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

func (r *specialRequestsRepository) GetSortedSubjects(ctx context.Context) ([]dto.SortedSubjectDTO, error) {
	sql := `
		SELECT subjects.name
		FROM public.subjects
		ORDER BY subjects.name ASC
	`

	log.Println("executing sql:", sql)

	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	var result []dto.SortedSubjectDTO
	for rows.Next() {
		var dto dto.SortedSubjectDTO
		err := rows.Scan(&dto.Name)
		if err != nil {
			return nil, handlePgError(err)
		}
		result = append(result, dto)
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *specialRequestsRepository) GetSortedMarks(ctx context.Context) ([]dto.SortedMarkDTO, error) {
	sql := `
		SELECT marks.student_id, marks.mark, marks.date
		FROM public.marks
		ORDER BY marks.date ASC, marks.mark DESC
	`

	log.Println("executing sql:", sql)

	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	var result []dto.SortedMarkDTO
	for rows.Next() {
		var dto dto.SortedMarkDTO
		err := rows.Scan(
			&dto.StudentID,
			&dto.Mark,
			&dto.Date,
		)
		if err != nil {
			return nil, handlePgError(err)
		}
		result = append(result, dto)
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *specialRequestsRepository) GetStudentGroupCombs(ctx context.Context) ([]dto.StudentGroupCombDTO, error) {
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

func (r *specialRequestsRepository) GetStudentsWithCurators(ctx context.Context) ([]dto.StudentCuratorDTO, error) {
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

func (r *specialRequestsRepository) GetCuratorsWithStudents(ctx context.Context) ([]dto.StudentCuratorDTO, error) {
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

func (r *specialRequestsRepository) GetAllStudentCuratorPairs(ctx context.Context) ([]dto.StudentCuratorDTO, error) {
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

func (r *specialRequestsRepository) GetStudentsUppercaseWithLength(ctx context.Context) ([]dto.StudentNameStatDTO, error) {
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

func (r *specialRequestsRepository) IsTeacher(ctx context.Context, id uint64) (dto.EmployeeRoleDTO, error) {
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
