package app

import (
	"github.com/adityaqudaedah/go_restful_api/controller"
	"github.com/adityaqudaedah/go_restful_api/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router{
	router := httprouter.New()

	router.POST("/api/categories",categoryController.Create)
	router.GET("/api/categories",categoryController.FindAll)
	router.GET("/api/categories/:categoryId",categoryController.FindById)
	router.PUT("/api/categories/:categoryId",categoryController.Update)
	router.DELETE("/api/categories/:categoryId",categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}