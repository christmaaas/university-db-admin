package postgres

import (
	"context"
	"log"
	"university-db-admin/internal/domain"
	"university-db-admin/internal/repository"

	"github.com/jackc/pgx/v5"
)

type groupsRepository struct {
	db *pgx.Conn
}

func NewGroupsRepository(db *pgx.Conn) repository.Groups {
	return &groupsRepository{
		db: db,
	}
}

func (g *groupsRepository) Create(ctx context.Context, grp domain.Group) error {
	sql := `
		INSERT INTO public.groups (number)
		VALUES ($1)
		RETURNING id
	`

	log.Println("executing sql: ", sql)
	err := g.db.QueryRow(ctx, sql, grp.Number).Scan(&grp.ID)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result: ", grp.ID)
	return nil
}

func (g *groupsRepository) FindOne(ctx context.Context, id uint64) (domain.Group, error) {
	sql := `
		SELECT g.id, g.number
		FROM public.groups g
		WHERE g.id = $1
	`

	var grp domain.Group

	log.Println("executing sql: ", sql)
	err := g.db.QueryRow(ctx, sql, id).Scan(&grp.ID, &grp.Number)
	if err != nil {
		return domain.Group{}, handlePgError(err)
	}

	log.Println("sql result: ", grp)
	return grp, nil
}

func (g *groupsRepository) FindAll(ctx context.Context) ([]domain.Group, error) {
	sql := `
		SELECT g.id, g.number 
		FROM public.groups g
	`

	var groups []domain.Group
	log.Println("executing sql: ", sql)

	rows, err := g.db.Query(ctx, sql)
	if err != nil {
		return groups, err
	}
	defer rows.Close()

	for rows.Next() {
		var grp domain.Group
		err := rows.Scan(&grp.ID, &grp.Number)
		if err != nil {
			return nil, handlePgError(err)
		}
		groups = append(groups, grp)
	}

	log.Println("sql result: ", groups)
	return groups, nil
}

func (g *groupsRepository) FindByNumber(ctx context.Context, num uint64) (domain.Group, error) {
	sql := `
		SELECT g.id, g.number
		FROM public.groups g
		WHERE g.number = $1
	`

	var grp domain.Group

	log.Println("executing sql: ", sql)
	err := g.db.QueryRow(ctx, sql, num).Scan(&grp.ID, &grp.Number)
	if err != nil {
		return domain.Group{}, handlePgError(err)
	}

	log.Println("sql result: ", grp)
	return grp, nil
}

func (g *groupsRepository) Update(ctx context.Context, id uint64, grp domain.Group) error {
	sql := `
		UPDATE public.groups
		SET number = $1
		WHERE id = $2
		RETURNING id
	`

	log.Println("executing sql: ", sql)
	err := g.db.QueryRow(ctx, sql, grp.Number, id).Scan(&id)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result: ", id)
	return nil
}

func (g *groupsRepository) Delete(ctx context.Context, id uint64) error {
	sql := `
		DELETE FROM public.groups
		WHERE id = $1
		RETURNING id
	`

	log.Println("executing sql: ", sql)
	err := g.db.QueryRow(ctx, sql, id).Scan(&id)
	if err != nil {
		return handlePgError(err)
	}

	log.Println("sql result: ", id)
	return nil
}
