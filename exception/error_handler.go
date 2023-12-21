package exception

import (
	"net/http"

	"github.com/adityaqudaedah/go_restful_api/helpers"
	"github.com/adityaqudaedah/go_restful_api/model/web"
	"github.com/go-playground/validator/v10"
)

func ErrorHandler(w http.ResponseWriter,req *http.Request,err interface{}) {

	if notFoundError(w,req,err) {
		return
	}
	if validatioError(w,req,err){
		return
	}
	internalServerError(w,req,err)
}

func internalServerError(w http.ResponseWriter,req *http.Request,err interface{})  {
	w.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code : http.StatusInternalServerError,
		Message: "INTERNAL SERVER ERROR",
		Status: "ERROR",
		Data: err,
	}
	helpers.WriteToResponseBody(w,webResponse)
}

func notFoundError(w http.ResponseWriter,req *http.Request,err interface{})bool{
	exception,ok := err.(NotFoundError)

	if ok {
		w.WriteHeader(http.StatusNotFound)

	webResponse := web.WebResponse{
		Code : http.StatusNotFound,
		Message: "NOT FOUND",
		Status: "ERROR",
		Data: exception.Error,
	}
	helpers.WriteToResponseBody(w,webResponse)
	return true
	}
	return false
}

func validatioError(w http.ResponseWriter,req *http.Request,err interface{})bool  {
	exception,ok := err.(validator.ValidationErrors)

	if ok {
		w.WriteHeader(http.StatusBadRequest)

	webResponse := web.WebResponse{
		Code : http.StatusBadRequest,
		Message: "BAD REQUEST",
		Status: "ERROR",
		Data: exception.Error(),
	}
	helpers.WriteToResponseBody(w,webResponse)
	return true
	}
	return false
}

