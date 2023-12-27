package main

import (
	"net/http"

	"github.com/adityaqudaedah/go_restful_api/app"
	"github.com/adityaqudaedah/go_restful_api/controller"
	"github.com/adityaqudaedah/go_restful_api/helpers"
	"github.com/adityaqudaedah/go_restful_api/middleware"
	"github.com/adityaqudaedah/go_restful_api/repository"
	"github.com/adityaqudaedah/go_restful_api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository,db,validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)
	server := http.Server{
		Addr: ":8080",Handler: middleware.NewAuthMiddleWare(router),
	}

	err := server.ListenAndServe()

	helpers.PanicIfError(err)
}