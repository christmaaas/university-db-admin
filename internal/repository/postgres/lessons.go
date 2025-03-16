package postgres

import (
	"context"
	"log"
	"university-db-admin/internal/domain"
	"university-db-admin/internal/repository"

	"github.com/jackc/pgx/v5"
)

type lessonsRepository struct {
	db *pgx.Conn
}

func NewLessonsRepository(db *pgx.Conn) repository.Lessons {
	return &lessonsRepository{
		db: db,
	}
}

func (l *lessonsRepository) Create(ctx context.Context, lsn domain.Lesson) error {
	sql := `
		INSERT INTO public.lessons (group_id, subject_id, lesson_type_id, week, weekday, room)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := l.db.QueryRow(ctx, sql,
		lsn.GroupID,
		lsn.SubjectID,
		lsn.LessonTypeID,
		lsn.Week,
		lsn.Weekday,
		lsn.Room,
	).Scan(&lsn.ID)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", lsn.ID)
	return nil
}

func (l *lessonsRepository) FindOne(ctx context.Context, id uint64) (domain.Lesson, error) {
	sql := `
		SELECT id, group_id, subject_id, lesson_type_id, week, weekday, room 
		FROM public.lessons
		WHERE id = $1
	`

	var lsn domain.Lesson
	log.Println("executing sql:", sql)
	err := l.db.QueryRow(ctx, sql, id).Scan(
		&lsn.ID,
		&lsn.GroupID,
		&lsn.SubjectID,
		&lsn.LessonTypeID,
		&lsn.Week,
		&lsn.Weekday,
		&lsn.Room,
	)
	if err != nil {
		return domain.Lesson{}, handlePgError(err)
	}

	log.Println("sql result:", lsn)
	return lsn, nil
}

func (l *lessonsRepository) FindAll(ctx context.Context) ([]domain.Lesson, error) {
	sql := `
		SELECT id, group_id, subject_id, lesson_type_id, week, weekday, room 
		FROM public.lessons
	`

	var lessons []domain.Lesson
	log.Println("executing sql:", sql)

	rows, err := l.db.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var lsn domain.Lesson
		err := rows.Scan(
			&lsn.ID,
			&lsn.GroupID,
			&lsn.SubjectID,
			&lsn.LessonTypeID,
			&lsn.Week,
			&lsn.Weekday,
			&lsn.Room,
		)
		if err != nil {
			return nil, handlePgError(err)
		}
		lessons = append(lessons, lsn)
	}

	log.Println("sql result:", lessons)
	return lessons, nil
}

func (l *lessonsRepository) findByField(ctx context.Context, field string, value interface{}) ([]domain.Lesson, error) {
	sql := `
		SELECT id, group_id, subject_id, lesson_type_id, week, weekday, room 
		FROM public.lessons
		WHERE ` + field + ` = $1
	`

	var lessons []domain.Lesson
	log.Println("executing sql:", sql)

	rows, err := l.db.Query(ctx, sql, value)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var lsn domain.Lesson
		err := rows.Scan(
			&lsn.ID,
			&lsn.GroupID,
			&lsn.SubjectID,
			&lsn.LessonTypeID,
			&lsn.Week,
			&lsn.Weekday,
			&lsn.Room,
		)
		if err != nil {
			return nil, handlePgError(err)
		}
		lessons = append(lessons, lsn)
	}

	log.Println("sql result:", lessons)
	return lessons, nil
}

func (l *lessonsRepository) FindByGroupID(ctx context.Context, id uint64) ([]domain.Lesson, error) {
	return l.findByField(ctx, "group_id", id)
}

func (l *lessonsRepository) FindBySubjectID(ctx context.Context, id uint64) ([]domain.Lesson, error) {
	return l.findByField(ctx, "subject_id", id)
}

func (l *lessonsRepository) FindByLessonTypeID(ctx context.Context, id uint64) ([]domain.Lesson, error) {
	return l.findByField(ctx, "lesson_type_id", id)
}

func (l *lessonsRepository) FindByWeek(ctx context.Context, week uint16) ([]domain.Lesson, error) {
	return l.findByField(ctx, "week", week)
}

func (l *lessonsRepository) FindByWeekday(ctx context.Context, weekday uint16) ([]domain.Lesson, error) {
	return l.findByField(ctx, "weekday", weekday)
}

func (l *lessonsRepository) FindByRoom(ctx context.Context, room uint64) ([]domain.Lesson, error) {
	return l.findByField(ctx, "room", room)
}

func (l *lessonsRepository) Update(ctx context.Context, id uint64, lsn domain.Lesson) error {
	sql := `
		UPDATE public.lessons
		SET group_id = $1, subject_id = $2, lesson_type_id = $3, week = $4, weekday = $5, room = $6
		WHERE id = $7
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := l.db.QueryRow(ctx, sql,
		lsn.GroupID,
		lsn.SubjectID,
		lsn.LessonTypeID,
		lsn.Week,
		lsn.Weekday,
		lsn.Room,
		id,
	).Scan(&id)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", id)
	return nil
}

func (l *lessonsRepository) Delete(ctx context.Context, id uint64) error {
	sql := `
		DELETE FROM public.lessons
		WHERE id = $1
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := l.db.QueryRow(ctx, sql, id).Scan(&id)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", id)
	return nil
}
