package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/adityaqudaedah/go_restful_api/helpers"
	"github.com/adityaqudaedah/go_restful_api/model/web"
	"github.com/adityaqudaedah/go_restful_api/service"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter,req *http.Request,params httprouter.Params) {
	decoder := json.NewDecoder(req.Body)
	var categoryCreateRequest =  web.CategoryCreateRequest{}
	err := decoder.Decode(&categoryCreateRequest)
	helpers.PanicIfError(err)

	categoryResponse := controller.CategoryService.Create(req.Context(),categoryCreateRequest)
	webResponse := web.WebResponse{
		Code: http.StatusOK, Message:"Successfuly added data" ,Status: "success",Data: categoryResponse,
	}

	w.Header().Add("Content-Type","application/json")
	encoder := json.NewEncoder(w)
	errEncoder := encoder.Encode(webResponse)

	helpers.PanicIfError(errEncoder)
}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter,req *http.Request,params httprouter.Params) {
	decoder := json.NewDecoder(req.Body)
	var categoryUpdateRequest = web.CategoryUpdateRequest{}
	err := decoder.Decode(&categoryUpdateRequest)
	helpers.PanicIfError(err)

	strCategoryId := params.ByName("categoryId")
	categoryId,errConv := strconv.Atoi(strCategoryId)
	helpers.PanicIfError(errConv)

 	categoryUpdateRequest.Id = categoryId

	categoryResponse := controller.CategoryService.Update(req.Context(),categoryUpdateRequest)

	

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Message: "Successfully updated",
		Status: "success",
		Data: categoryResponse,
	}
	w.Header().Add("Content-Type","application/json")
	encoder := json.NewEncoder(w)
	errEncoder := encoder.Encode(webResponse)

	helpers.PanicIfError(errEncoder)
}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter,req *http.Request,params httprouter.Params) {	
	strCategoryId := params.ByName("categoryId")
	categoryId,errConv := strconv.Atoi(strCategoryId)
	helpers.PanicIfError(errConv)
	
	controller.CategoryService.Delete(req.Context(),categoryId)

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Message: "Successfully delete",
		Status: "success",
	}

	w.Header().Add("Content-Type","application/json")
	encoder := json.NewEncoder(w)
	errEncoder := encoder.Encode(webResponse)

	helpers.PanicIfError(errEncoder)
}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter,req *http.Request,params httprouter.Params) {
	strCategoryId := params.ByName("categoryId")
	categoryId,errConv := strconv.Atoi(strCategoryId)

	helpers.PanicIfError(errConv)

	categoryResponse := controller.CategoryService.FindById(req.Context(),categoryId)

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Message: "Successfully find data by id",
		Status: "success",
		Data: categoryResponse,
	}

	w.Header().Add("Content-Type","application/json")
	encoder := json.NewEncoder(w)
	errEncoder := encoder.Encode(webResponse)

	helpers.PanicIfError(errEncoder)
}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter,req *http.Request,params httprouter.Params) {
	categoryResponse := controller.CategoryService.FindAll(req.Context())

	webResponse := web.WebResponse{
		Code: http.StatusOK,
		Message: "Successfully get all data",
		Status: "success",
		Data: categoryResponse,
	}

	w.Header().Add("Content-Type","application/json")
	encoder := json.NewEncoder(w)
	errEncoder := encoder.Encode(webResponse)

	helpers.PanicIfError(errEncoder)
}