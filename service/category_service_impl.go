package service

import (
	"context"
	"database/sql"

	"github.com/adityaqudaedah/go_restful_api/exception"
	"github.com/adityaqudaedah/go_restful_api/helpers"
	"github.com/adityaqudaedah/go_restful_api/model/domain"
	"github.com/adityaqudaedah/go_restful_api/model/web"
	"github.com/adityaqudaedah/go_restful_api/repository"
	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB *sql.DB
	Validator *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository,db *sql.DB,validator *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB: db,
		Validator: validator,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse  {
	errValidate := service.Validator.Struct(request)
	helpers.PanicIfError(errValidate)


	tx,err := service.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	helpers.PanicIfError(err)

	category := domain.Category{Name: request.Name}

	 category = service.CategoryRepository.Create(ctx,tx,category)

	categoryResponse,errFindById := service.CategoryRepository.FindById(ctx,tx,category.Id)

	helpers.PanicIfError(errFindById)
	
	return helpers.ToCategoryResponse(categoryResponse)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse{
	errValidate := service.Validator.Struct(request)
	helpers.PanicIfError(errValidate)

    tx,err := service.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	helpers.PanicIfError(err)

	category,errFindById := service.CategoryRepository.FindById(ctx,tx,request.Id)
	
	if errFindById != nil{
		panic(exception.NewNotFoundError("not found"))
	}

	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx,tx,category) 

	
	return helpers.ToCategoryResponse(category)
}

func(service *CategoryServiceImpl) Delete(ctx context.Context,requestId int){
	tx,err := service.DB.Begin()
	defer helpers.CommitOrRollback(tx)

	helpers.PanicIfError(err)

	category,errFindById := service.CategoryRepository.FindById(ctx,tx,requestId)
	
	if errFindById != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}


	service.CategoryRepository.Delete(ctx,tx,category.Id)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context,requestId int) web.CategoryResponse{
	tx,err := service.DB.Begin()

	helpers.PanicIfError(err)

	category,errFindById := service.CategoryRepository.FindById(ctx,tx,requestId)

	if errFindById != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helpers.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse{
	tx,err := service.DB.Begin()

	helpers.PanicIfError(err)

	categories := service.CategoryRepository.FindAll(ctx,tx)

	return helpers.ToCategoryResponses(categories...)
}