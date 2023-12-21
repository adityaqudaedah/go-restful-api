package main

import (
	"net/http"

	"github.com/adityaqudaedah/go_restful_api/app"
	"github.com/adityaqudaedah/go_restful_api/controller"
	"github.com/adityaqudaedah/go_restful_api/exception"
	"github.com/adityaqudaedah/go_restful_api/helpers"
	"github.com/adityaqudaedah/go_restful_api/middleware"
	"github.com/adityaqudaedah/go_restful_api/repository"
	"github.com/adityaqudaedah/go_restful_api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository,db,validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.POST("/api/categories",categoryController.Create)
	router.GET("/api/categories",categoryController.FindAll)
	router.GET("/api/categories/:categoryId",categoryController.FindById)
	router.PUT("/api/categories/:categoryId",categoryController.Update)
	router.DELETE("/api/categories/:categoryId",categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr: ":8080",Handler: middleware.NewAuthMiddleWare(router),
	}

	err := server.ListenAndServe()

	helpers.PanicIfError(err)
}