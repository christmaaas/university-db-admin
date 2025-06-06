package postgres

import (
	"context"
	"log"
	"university-db-admin/internal/domain"
	"university-db-admin/internal/dto"
	"university-db-admin/internal/repository"

	"github.com/jackc/pgx/v5"
)

type subjectsRepository struct {
	db *pgx.Conn
}

func NewSubjectsRepository(db *pgx.Conn) repository.Subjects {
	return &subjectsRepository{
		db: db,
	}
}

func (s *subjectsRepository) Create(ctx context.Context, sbj domain.Subject) error {
	sql := `
		INSERT INTO public.subjects (name, description)
		VALUES ($1, $2)
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := s.db.QueryRow(ctx, sql, sbj.Name, sbj.Description).Scan(&sbj.ID)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", sbj.ID)
	return nil
}

func (s *subjectsRepository) FindOne(ctx context.Context, id uint64) (domain.Subject, error) {
	sql := `
		SELECT id, name, description
		FROM public.subjects
		WHERE id = $1
	`

	var sbj domain.Subject
	log.Println("executing sql:", sql)
	err := s.db.QueryRow(ctx, sql, id).Scan(&sbj.ID, &sbj.Name, &sbj.Description)
	if err != nil {
		return domain.Subject{}, handlePgError(err)
	}

	log.Println("sql result:", sbj)
	return sbj, nil
}

func (s *subjectsRepository) FindAll(ctx context.Context) ([]domain.Subject, error) {
	sql := `
		SELECT id, name, description
		FROM public.subjects
	`

	var subjects []domain.Subject
	log.Println("executing sql:", sql)

	rows, err := s.db.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var sbj domain.Subject
		err := rows.Scan(&sbj.ID, &sbj.Name, &sbj.Description)
		if err != nil {
			return nil, handlePgError(err)
		}
		subjects = append(subjects, sbj)
	}

	log.Println("sql result:", subjects)
	return subjects, nil
}

func (s *subjectsRepository) FindByName(ctx context.Context, name string) (domain.Subject, error) {
	sql := `
		SELECT id, name, description
		FROM public.subjects
		WHERE name = $1
	`

	var sbj domain.Subject
	log.Println("executing sql:", sql)

	err := s.db.QueryRow(ctx, sql, name).Scan(&sbj.ID, &sbj.Name, &sbj.Description)
	if err != nil {
		return domain.Subject{}, handlePgError(err)
	}

	log.Println("sql result:", sbj)
	return sbj, nil
}

func (r *subjectsRepository) FindAllSorted(ctx context.Context) ([]dto.SortedSubjectDTO, error) {
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

func (s *subjectsRepository) Update(ctx context.Context, id uint64, sbj domain.Subject) error {
	sql := `
		UPDATE public.subjects
		SET name = $1, description = $2
		WHERE id = $3
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := s.db.QueryRow(ctx, sql, sbj.Name, sbj.Description, id).Scan(&id)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", id)
	return nil
}

func (s *subjectsRepository) Delete(ctx context.Context, id uint64) error {
	sql := `
		DELETE FROM public.subjects
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
