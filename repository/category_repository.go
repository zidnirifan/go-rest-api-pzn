package repository

import (
	"context"
	"database/sql"
	"go-restful/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, categoryId int)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
}
