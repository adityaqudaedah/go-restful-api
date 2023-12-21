package repository

import (
	"context"
	"database/sql"

	"github.com/adityaqudaedah/go_restful_api/model/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context,tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx,category domain.Category) domain.Category
	Delete(ctx context.Context,tx *sql.Tx,categoryId int)
	FindById(ctx context.Context, tx *sql.Tx,categoryId int) (domain.Category,error)
	FindAll(ctx context.Context, tx *sql.Tx,) []domain.Category
}