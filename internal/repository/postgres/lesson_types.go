package postgres

import (
	"context"
	"log"
	"university-db-admin/internal/domain"
	"university-db-admin/internal/repository"

	"github.com/jackc/pgx/v5"
)

type lessonTypesRepository struct {
	dbclient *pgx.Conn
}

func NewLessonTypesRepository(dbclient *pgx.Conn) repository.LessonTypes {
	return &lessonTypesRepository{
		dbclient: dbclient,
	}
}

func (l *lessonTypesRepository) Create(ctx context.Context, lsn domain.LessonType) error {
	sql := `
		INSERT INTO public.lesson_types (name)
		VALUES ($1)
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := l.dbclient.QueryRow(ctx, sql, lsn.Name).Scan(&lsn.ID)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", lsn.ID)
	return nil
}

func (l *lessonTypesRepository) FindOne(ctx context.Context, id uint64) (domain.LessonType, error) {
	sql := `
		SELECT lt.id, lt.name 
		FROM public.lesson_types lt
		WHERE lt.id = $1
	`

	var lsn domain.LessonType
	log.Println("executing sql:", sql)
	err := l.dbclient.QueryRow(ctx, sql, id).Scan(&lsn.ID, &lsn.Name)
	if err != nil {
		return domain.LessonType{}, handlePgError(err)
	}

	log.Println("sql result:", lsn)
	return lsn, nil
}

func (l *lessonTypesRepository) FindAll(ctx context.Context) ([]domain.LessonType, error) {
	sql := `
		SELECT lt.id, lt.name 
		FROM public.lesson_types lt
	`

	var lessonTypes []domain.LessonType
	log.Println("executing sql:", sql)

	rows, err := l.dbclient.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var lsn domain.LessonType
		err := rows.Scan(&lsn.ID, &lsn.Name)
		if err != nil {
			return nil, handlePgError(err)
		}
		lessonTypes = append(lessonTypes, lsn)
	}

	log.Println("sql result:", lessonTypes)
	return lessonTypes, nil
}

func (l *lessonTypesRepository) FindByName(ctx context.Context, name string) (domain.LessonType, error) {
	sql := `
		SELECT lt.id, lt.name 
		FROM public.lesson_types lt
		WHERE lt.name = $1
	`

	var lsn domain.LessonType
	log.Println("executing sql:", sql)
	err := l.dbclient.QueryRow(ctx, sql, name).Scan(&lsn.ID, &lsn.Name)
	if err != nil {
		return domain.LessonType{}, handlePgError(err)
	}

	log.Println("sql result:", lsn)
	return lsn, nil
}

func (l *lessonTypesRepository) Update(ctx context.Context, id uint64, lsn domain.LessonType) error {
	sql := `
		UPDATE public.lesson_types
		SET name = $1
		WHERE id = $2
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := l.dbclient.QueryRow(ctx, sql, lsn.Name, id).Scan(&id)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", id)
	return nil
}

func (l *lessonTypesRepository) Delete(ctx context.Context, id uint64) error {
	sql := `
		DELETE FROM public.lesson_types
		WHERE id = $1
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := l.dbclient.QueryRow(ctx, sql, id).Scan(&id)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", id)
	return nil
}
