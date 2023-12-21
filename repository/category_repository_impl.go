package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/adityaqudaedah/go_restful_api/helpers"
	"github.com/adityaqudaedah/go_restful_api/model/domain"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Create(ctx context.Context,tx *sql.Tx,category domain.Category) domain.Category {
	SQL := "INSERT INTO category (name) values(?)"
	result,errSql := tx.ExecContext(ctx,SQL,category.Name)

	helpers.PanicIfError(errSql)

	id ,err:= result.LastInsertId()

	helpers.PanicIfError(err)

	category.Id = int(id)

	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category)domain.Category  {
	SQL := "UPDATE category SET name = ? WHERE id = ?"
	_,errSql := tx.ExecContext(ctx,SQL,category.Name,category.Id)

	fmt.Println("================",errSql)

	helpers.PanicIfError(errSql)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, categoryId int)  {
	SQL := "DELETE FROM category WHERE id = ?"
	_,errSql := tx.ExecContext(ctx,SQL,categoryId)

	helpers.PanicIfError(errSql)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx,categoryId int) (domain.Category,error) {
	SQL := "SELECT id, name, created_at FROM category WHERE id = ?"
	rows,errSql := tx.QueryContext(ctx,SQL,categoryId)
	helpers.PanicIfError(errSql)
	defer rows.Close()
	category := domain.Category{

	}
	if rows.Next(){
		err := rows.Scan(&category.Id,&category.Name,&category.CreatedAt)
		helpers.PanicIfError(err)
		return category,nil
	}else{
		return category,errors.New("category not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context,tx *sql.Tx) []domain.Category {
	SQL := "SELECT id,name,created_at FROM category"
	rows,errSql := tx.QueryContext(ctx,SQL)
	helpers.PanicIfError(errSql)
	defer rows.Close()
	var categories []domain.Category

	for rows.Next(){
		category := domain.Category{}
		err:= rows.Scan(&category.Id,&category.Name,&category.CreatedAt)
		helpers.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}

