package postgres

import (
	"context"
	"fmt"
	"log"
	"university-db-admin/internal/repository"

	"github.com/jackc/pgx/v5"
)

type specialRequestsRepository struct {
	dbclient *pgx.Conn
}

func NewSpecialRepository(dbclient *pgx.Conn) repository.Special {
	return &specialRequestsRepository{
		dbclient: dbclient,
	}
}

func (r *specialRequestsRepository) IsEmployeeTeacher(ctx context.Context, id uint64) (bool, error) {
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

func (r *specialRequestsRepository) GetScheduleByGroups(ctx context.Context) ([][]string, error) {
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
			room        string
			week        int
			weekday     int
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
			room,
			fmt.Sprintf("%d", week),
			fmt.Sprintf("%d", weekday),
		})
	}

	log.Println("sql result:", result)
	return result, nil
}

func (r *specialRequestsRepository) GetAllEmployeesInfo(ctx context.Context) ([][]string, error) {
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
