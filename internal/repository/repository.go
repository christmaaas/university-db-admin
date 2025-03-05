package repository

import (
	"context"
	"university-db-admin/internal/domain"
)

type Employees interface {
	Create(ctx context.Context, emp domain.Employee) error
	FindOne(ctx context.Context, id uint64) (domain.Employee, error)
	FindAll(ctx context.Context) ([]domain.Employee, error)
	FindByName(ctx context.Context, name string) ([]domain.Employee, error)
	FindByPassport(ctx context.Context, passport string) (domain.Employee, error)
	FindByPosition(ctx context.Context, position uint64) ([]domain.Employee, error)
	Update(ctx context.Context, id uint64, emp domain.Employee) error
	Delete(ctx context.Context, id uint64) error
}
