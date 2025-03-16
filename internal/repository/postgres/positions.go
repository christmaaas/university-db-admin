package postgres

import (
	"context"
	"log"
	"university-db-admin/internal/domain"
	"university-db-admin/internal/repository"

	"github.com/jackc/pgx/v5"
)

type positionsRepository struct {
	db *pgx.Conn
}

func NewPositionsRepository(db *pgx.Conn) repository.Positions {
	return &positionsRepository{
		db: db,
	}
}

func (p *positionsRepository) Create(ctx context.Context, pos domain.Position) error {
	sql := `
		INSERT INTO public.positions (name)
		VALUES ($1)
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := p.db.QueryRow(ctx, sql, pos.Name).Scan(&pos.ID)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", pos.ID)
	return nil
}

func (p *positionsRepository) FindOne(ctx context.Context, id uint64) (domain.Position, error) {
	sql := `
		SELECT id, name
		FROM public.positions
		WHERE id = $1
	`

	var pos domain.Position
	log.Println("executing sql:", sql)
	err := p.db.QueryRow(ctx, sql, id).Scan(&pos.ID, &pos.Name)
	if err != nil {
		return domain.Position{}, handlePgError(err)
	}

	log.Println("sql result:", pos)
	return pos, nil
}

func (p *positionsRepository) FindAll(ctx context.Context) ([]domain.Position, error) {
	sql := `
		SELECT id, name
		FROM public.positions
	`

	var positions []domain.Position
	log.Println("executing sql:", sql)

	rows, err := p.db.Query(ctx, sql)
	if err != nil {
		return nil, handlePgError(err)
	}
	defer rows.Close()

	for rows.Next() {
		var pos domain.Position
		err := rows.Scan(&pos.ID, &pos.Name)
		if err != nil {
			return nil, handlePgError(err)
		}
		positions = append(positions, pos)
	}

	log.Println("sql result:", positions)
	return positions, nil
}

func (p *positionsRepository) FindByName(ctx context.Context, name string) (domain.Position, error) {
	sql := `
		SELECT id, name
		FROM public.positions
		WHERE name = $1
	`

	var pos domain.Position

	log.Println("executing sql:", sql)
	err := p.db.QueryRow(ctx, sql, name).Scan(&pos.ID, &pos.Name)
	if err != nil {
		return domain.Position{}, handlePgError(err)
	}

	log.Println("sql result:", pos)
	return pos, nil
}

func (p *positionsRepository) Update(ctx context.Context, id uint64, pos domain.Position) error {
	sql := `
		UPDATE public.positions
		SET name = $1
		WHERE id = $2
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := p.db.QueryRow(ctx, sql, pos.Name, id).Scan(&id)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", id)
	return nil
}

func (p *positionsRepository) Delete(ctx context.Context, id uint64) error {
	sql := `
		DELETE FROM public.positions
		WHERE id = $1
		RETURNING id
	`

	log.Println("executing sql:", sql)
	err := p.db.QueryRow(ctx, sql, id).Scan(&id)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result:", id)
	return nil
}
