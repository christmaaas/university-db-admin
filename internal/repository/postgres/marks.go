package postgres

import (
	"context"
	"log"
	"time"
	"university-db-admin/internal/domain"
	"university-db-admin/internal/repository"

	"github.com/jackc/pgx/v5"
)

type marksRepository struct {
	db *pgx.Conn
}

func NewMarksRepository(db *pgx.Conn) repository.Marks {
	return &marksRepository{
		db: db,
	}
}

func (m *marksRepository) Create(ctx context.Context, mark domain.Mark) error {
	sql := `
		INSERT INTO public.marks (employee_id, student_id, subject_id, mark, date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := m.db.QueryRow(ctx, sql,
		mark.EmployeeID,
		mark.StudentID,
		mark.SubjectID,
		mark.Mark,
		mark.Date,
	).Scan(&mark.ID)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", mark.ID)
	return nil
}

func (m *marksRepository) FindOne(ctx context.Context, id uint64) (domain.Mark, error) {
	sql := `
        SELECT id, employee_id, student_id, subject_id, mark, date 
        FROM public.marks
        WHERE id = $1
    `

	var (
		mark domain.Mark
		date time.Time
	)

	log.Println("executing sql:", sql)
	err := m.db.QueryRow(ctx, sql, id).Scan(
		&mark.ID,
		&mark.EmployeeID,
		&mark.StudentID,
		&mark.SubjectID,
		&mark.Mark,
		&date,
	)
	if err != nil {
		return domain.Mark{}, handlePgError(err)
	}

	log.Println("sql result:", mark)
	mark.Date = date.Format("2006-01-02")
	return mark, nil
}

func (m *marksRepository) FindAll(ctx context.Context) ([]domain.Mark, error) {
	sql := `
		SELECT id, employee_id, student_id, subject_id, mark, date 
		FROM public.marks
	`

	var marks []domain.Mark
	log.Println("executing sql:", sql)

	rows, err := m.db.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			mark domain.Mark
			date time.Time
		)
		err := rows.Scan(
			&mark.ID,
			&mark.EmployeeID,
			&mark.StudentID,
			&mark.SubjectID,
			&mark.Mark,
			&date,
		)
		if err != nil {
			return nil, handlePgError(err)
		}

		mark.Date = date.Format("2006-01-02")
		marks = append(marks, mark)
	}

	log.Println("sql result:", marks)
	return marks, nil
}

func (m *marksRepository) findByField(ctx context.Context, field string, value interface{}) ([]domain.Mark, error) {
	sql := `
		SELECT id, employee_id, student_id, subject_id, mark, date 
		FROM public.marks
		WHERE ` + field + ` = $1
	`

	var marks []domain.Mark
	log.Println("executing sql:", sql)

	rows, err := m.db.Query(ctx, sql, value)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			mark domain.Mark
			date time.Time
		)
		err := rows.Scan(
			&mark.ID,
			&mark.EmployeeID,
			&mark.StudentID,
			&mark.SubjectID,
			&mark.Mark,
			&date,
		)
		if err != nil {
			return nil, handlePgError(err)
		}

		mark.Date = date.Format("2006-01-02")
		marks = append(marks, mark)
	}

	log.Println("sql result:", marks)
	return marks, nil
}

func (m *marksRepository) FindByEmployeeID(ctx context.Context, id uint64) ([]domain.Mark, error) {
	return m.findByField(ctx, "employee_id", id)
}

func (m *marksRepository) FindByStudentID(ctx context.Context, id uint64) ([]domain.Mark, error) {
	return m.findByField(ctx, "student_id", id)
}

func (m *marksRepository) FindBySubjectID(ctx context.Context, id uint64) ([]domain.Mark, error) {
	return m.findByField(ctx, "subject_id", id)
}

func (m *marksRepository) FindByMark(ctx context.Context, mark uint16) ([]domain.Mark, error) {
	return m.findByField(ctx, "mark", mark)
}

func (m *marksRepository) FindByDate(ctx context.Context, date string) ([]domain.Mark, error) {
	return m.findByField(ctx, "date", date)
}

func (m *marksRepository) Update(ctx context.Context, id uint64, mark domain.Mark) error {
	sql := `
		UPDATE public.marks
		SET employee_id = $1, student_id = $2, subject_id = $3, mark = $4, date = $5
		WHERE id = $6
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := m.db.QueryRow(ctx, sql,
		mark.EmployeeID,
		mark.StudentID,
		mark.SubjectID,
		mark.Mark,
		mark.Date,
		id,
	).Scan(&id)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", id)
	return nil
}

func (m *marksRepository) Delete(ctx context.Context, id uint64) error {
	sql := `
		DELETE FROM public.marks
		WHERE id = $1
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := m.db.QueryRow(ctx, sql, id).Scan(&id)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", id)
	return nil
}
