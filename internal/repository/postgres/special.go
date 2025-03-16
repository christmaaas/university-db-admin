package postgres

import (
	"context"
	"fmt"
	"log"
	"time"
	"university-db-admin/internal/repository"

	"github.com/jackc/pgx/v5"
)

type specialRequestsRepository struct {
	dbclient *pgx.Conn
}

func NewSpecialRequestsRepository(dbclient *pgx.Conn) repository.Special {
	return &specialRequestsRepository{
		dbclient: dbclient,
	}
}

func (r *specialRequestsRepository) GetAllEmployees(ctx context.Context) ([][]string, error) {
	sql := `
		SELECT employees.name, employees.passport
		FROM public.employees
	`

	var result [][]string
	log.Println("executing sql:", sql)

	rows, err := r.dbclient.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			name     string
			passport string
		)
		err := rows.Scan(
			&name,
			&passport,
		)
		if err != nil {
			return nil, handlePgError(err)
		}

		result = append(result, []string{
			name,
			passport,
		})
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *specialRequestsRepository) GetEmployeeByID(ctx context.Context, id uint64) ([][]string, error) {
	sql := `
		SELECT employees.name, employees.passport
		FROM public.employees
		WHERE employees.id = $1
	`

	var result [][]string
	log.Println("executing sql:", sql)

	rows, err := r.dbclient.Query(ctx, sql, id)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			name     string
			passport string
		)
		err := rows.Scan(
			&name,
			&passport,
		)
		if err != nil {
			return nil, handlePgError(err)
		}

		result = append(result, []string{
			name,
			passport,
		})
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *specialRequestsRepository) GetStudentsNoCurator(ctx context.Context) ([][]string, error) {
	sql := `
		SELECT students.name, 
			students.passport,
			students.group_id
		FROM public.students
		WHERE NOT students.employee_id IS NOT NULL
	`

	var result [][]string
	log.Println("executing sql:", sql)

	rows, err := r.dbclient.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			name     string
			passport string
			groupID  uint64
		)
		err := rows.Scan(
			&name,
			&passport,
			&groupID,
		)
		if err != nil {
			return nil, handlePgError(err)
		}

		result = append(result, []string{
			name,
			passport,
			fmt.Sprintf("%d", groupID),
		})
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *specialRequestsRepository) GetLessonsSchedule(ctx context.Context) ([][]string, error) {
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

	var result [][]string
	log.Println("executing sql:", sql)

	rows, err := r.dbclient.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			groupNumber string
			subjectName string
			lessonType  string
			room        uint64
			week        uint16
			weekday     uint16
		)
		err := rows.Scan(
			&groupNumber,
			&subjectName,
			&lessonType,
			&room,
			&week,
			&weekday,
		)
		if err != nil {
			return nil, handlePgError(err)
		}

		result = append(result, []string{
			groupNumber,
			subjectName,
			lessonType,
			fmt.Sprintf("%d", room),
			fmt.Sprintf("%d", week),
			fmt.Sprintf("%d", weekday),
		})
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *specialRequestsRepository) GetEmployeesByPositions(ctx context.Context, firstID, secondID uint64) ([][]string, error) {
	sql := `
		SELECT employees.name
		FROM public.employees
		WHERE employees.position_id = $1 OR employees.position_id = $2
	`

	var result [][]string
	log.Println("executing sql:", sql)

	rows, err := r.dbclient.Query(ctx, sql, firstID, secondID)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, handlePgError(err)
		}

		result = append(result, []string{name})
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *specialRequestsRepository) GetMarksBySubject(ctx context.Context, id uint64, m uint16) ([][]string, error) {
	sql := `
		SELECT marks.student_id, marks.mark, marks.date
		FROM public.marks
		WHERE subject_id = $1 AND mark > $2
	`

	var result [][]string
	log.Println("executing sql:", sql)

	rows, err := r.dbclient.Query(ctx, sql, id, m)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			studentID uint64
			mark      uint16
			date      time.Time
		)
		err := rows.Scan(
			&studentID,
			&mark,
			&date,
		)
		if err != nil {
			return nil, handlePgError(err)
		}

		result = append(result, []string{
			fmt.Sprintf("%d", studentID),
			fmt.Sprintf("%d", mark),
			date.Format("2006-01-02"),
		})
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *specialRequestsRepository) GetStudentsByMiddlename(ctx context.Context, m string) ([][]string, error) {
	sql := `
		SELECT students.name, students.passport
		FROM public.students
		WHERE students.name LIKE $1
	`

	var result [][]string
	log.Println("executing sql:", sql)

	pattern := "%" + m
	rows, err := r.dbclient.Query(ctx, sql, pattern)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			name     string
			passport string
		)
		err := rows.Scan(
			&name,
			&passport,
		)
		if err != nil {
			return nil, handlePgError(err)
		}

		result = append(result, []string{
			name,
			passport,
		})
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *specialRequestsRepository) GetSortedSubjects(ctx context.Context) ([][]string, error) {
	sql := `
		SELECT subjects.name
		FROM public.subjects
		ORDER BY subjects.name ASC
	`

	var result [][]string
	log.Println("executing sql:", sql)

	rows, err := r.dbclient.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, handlePgError(err)
		}

		result = append(result, []string{name})
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *specialRequestsRepository) GetSortedMarks(ctx context.Context) ([][]string, error) {
	sql := `
		SELECT marks.student_id, marks.mark, marks.date
		FROM public.marks
		ORDER BY marks.date ASC, marks.mark DESC
	`

	var result [][]string
	log.Println("executing sql:", sql)

	rows, err := r.dbclient.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			studentID uint64
			mark      uint16
			date      time.Time
		)
		err := rows.Scan(
			&studentID,
			&mark,
			&date,
		)
		if err != nil {
			return nil, handlePgError(err)
		}

		result = append(result, []string{
			fmt.Sprintf("%d", studentID),
			fmt.Sprintf("%d", mark),
			date.Format("2006-01-02"),
		})
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *specialRequestsRepository) GetStudentGroupCombs(ctx context.Context) ([][]string, error) {
	sql := `
		SELECT students.name, groups.number
		FROM public.students
		CROSS JOIN public.groups
	`

	var result [][]string
	log.Println("executing sql:", sql)

	rows, err := r.dbclient.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			name   string
			number uint64
		)
		err := rows.Scan(
			&name,
			&number,
		)
		if err != nil {
			return nil, handlePgError(err)
		}

		result = append(result, []string{
			name,
			fmt.Sprintf("%d", number),
		})
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *specialRequestsRepository) GetStudentsWithCurators(ctx context.Context) ([][]string, error) {
	sql := `
		SELECT students.name,
			students.passport,
			employees.name,
			employees.passport
		FROM public.students
		LEFT OUTER JOIN employees ON students.employee_id = employees.id;
	`

	var result [][]string
	log.Println("executing sql:", sql)

	rows, err := r.dbclient.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			studName     string
			studPassport string
			empName      string
			empPassport  string
		)
		err := rows.Scan(
			&studName,
			&studPassport,
			&empName,
			&empPassport,
		)
		if err != nil {
			return nil, handlePgError(err)
		}

		result = append(result, []string{
			studName,
			studPassport,
			empName,
			empPassport,
		})
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *specialRequestsRepository) GetCuratorsWithStudents(ctx context.Context) ([][]string, error) {
	sql := `
		SELECT students.name,
			students.passport,
			employees.name,
			employees.passport
		FROM public.students
		RIGHT OUTER JOIN employees ON students.employee_id = employees.id;
	`

	var result [][]string
	log.Println("executing sql:", sql)

	rows, err := r.dbclient.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			studName     string
			studPassport string
			empName      string
			empPassport  string
		)
		err := rows.Scan(
			&studName,
			&studPassport,
			&empName,
			&empPassport,
		)
		if err != nil {
			return nil, handlePgError(err)
		}

		result = append(result, []string{
			studName,
			studPassport,
			empName,
			empPassport,
		})
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *specialRequestsRepository) GetAllStudentCuratorPairs(ctx context.Context) ([][]string, error) {
	sql := `
		SELECT students.name,
			students.passport,
			employees.name,
			employees.passport
		FROM public.students
		FULL OUTER JOIN employees ON students.employee_id = employees.id;
	`

	var result [][]string
	log.Println("executing sql:", sql)

	rows, err := r.dbclient.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			studName     string
			studPassport string
			empName      string
			empPassport  string
		)
		err := rows.Scan(
			&studName,
			&studPassport,
			&empName,
			&empPassport,
		)
		if err != nil {
			return nil, handlePgError(err)
		}

		result = append(result, []string{
			studName,
			studPassport,
			empName,
			empPassport,
		})
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *specialRequestsRepository) GetStudentsUppercaseWithLength(ctx context.Context) ([][]string, error) {
	sql := `
		SELECT students.id,
			UPPER(students.name),
			LENGTH(students.name)
		FROM public.students
	`

	var result [][]string
	log.Println("executing sql:", sql)

	rows, err := r.dbclient.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id        uint64
			nameUpper string
			nameLen   uint64
		)
		err := rows.Scan(
			&id,
			&nameUpper,
			&nameLen,
		)
		if err != nil {
			return nil, handlePgError(err)
		}

		result = append(result, []string{
			fmt.Sprintf("%d", id),
			nameUpper,
			fmt.Sprintf("%d", nameLen),
		})
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *specialRequestsRepository) IsTeacher(ctx context.Context, id uint64) (bool, error) {
	const teacherName = "Преподаватель"
	query := `
		SELECT EXISTS (
			SELECT 1 
			FROM employees e
			INNER JOIN positions p ON e.position_id = p.id
			WHERE e.id = $1 AND p.name = $2
		)
	`

	var exists bool
	err := r.dbclient.QueryRow(ctx, query, id, teacherName).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
