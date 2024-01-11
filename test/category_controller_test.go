package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/adityaqudaedah/go_restful_api/app"
	"github.com/adityaqudaedah/go_restful_api/controller"
	"github.com/adityaqudaedah/go_restful_api/helpers"
	"github.com/adityaqudaedah/go_restful_api/middleware"
	"github.com/adityaqudaedah/go_restful_api/model/domain"
	"github.com/adityaqudaedah/go_restful_api/model/web"
	"github.com/adityaqudaedah/go_restful_api/repository"
	"github.com/adityaqudaedah/go_restful_api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)


var URL_ENDPOINT = "http://localhost:8080/api/categories"

var jsonBlob = strings.NewReader(`{"name" : "testifying"}`)

func dbConf() *sql.DB  {
	db,errDb := sql.Open("mysql","root@tcp(localhost:3306)/go_restful_api_test?parseTime=true")

	helpers.PanicIfError(errDb)

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60*time.Minute)

	return db
}

func truncateCategory(DB *sql.DB)  {
	DB.Exec("TRUNCATE category")
}

func routerConf(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository,db,validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleWare(router)
}
func TestCreateCategorySuccess(t *testing.T) {
	db := dbConf()
	truncateCategory(db)
	router := routerConf(db)

	stringReader := strings.NewReader(`{"name" : "testify"}`)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost,URL_ENDPOINT,stringReader)
	request.Header.Add("Content-Type","application/json")
	request.Header.Add("MAMAT-API-KEY","mamat-rahmat")

	router.ServeHTTP(recorder,request)

	response := recorder.Result()

	var responseWeb web.WebResponse

	decoder := json.NewDecoder(response.Body)

	errDecode := decoder.Decode(&responseWeb)

	helpers.PanicIfError(errDecode)

	assert.Equal(t,200,responseWeb.Code)
	assert.Equal(t,"testify",responseWeb.Data.(map[string]interface{})["name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	db := dbConf()
	router := routerConf(db)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost,URL_ENDPOINT,strings.NewReader(`{"name" : " "}`))
	request.Header.Add("Content-Type","application/json")
	request.Header.Add("MAMAT-API-KEY","mamat-rahmat")
	
	router.ServeHTTP(recorder,request)

	r := recorder.Result().Body

	decoder := json.NewDecoder(r)

	var responseWeb web.WebResponse

	errDecode := decoder.Decode(&responseWeb)

	helpers.PanicIfError(errDecode)

	assert.Equal(t,400,responseWeb.Code)
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := dbConf()
	truncateCategory(db)
	router := routerConf(db)

	tx,_ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(),tx,domain.Category{
		Name : "test_craete",
	})
	tx.Commit()

	recorder:= httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPut,URL_ENDPOINT + "/" + strconv.Itoa(category.Id),jsonBlob)

	request.Header.Add("Content-Type","application/json")
	request.Header.Add("MAMAT-API-KEY","mamat-rahmat")

	router.ServeHTTP(recorder,request)

	r := recorder.Result().Body

	var responseWeb web.WebResponse

	decoder := json.NewDecoder(r)
	errDecode := decoder.Decode(&responseWeb)
	helpers.PanicIfError(errDecode)

	assert.Equal(t,"testifying",responseWeb.Data.(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db:= dbConf()
	router := routerConf(db)

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPut,URL_ENDPOINT + "/" + "10",jsonBlob)
	request.Header.Add("Content-Type","application/json")
	request.Header.Add("MAMAT-API-KEY","mamat-rahmat")

	router.ServeHTTP(recorder,request)

	r := recorder.Result().Body
	var responseWeb web.WebResponse
	decoder := json.NewDecoder(r)
	errDecode := decoder.Decode(&responseWeb)

	helpers.PanicIfError(errDecode)


	assert.Equal(t,http.StatusBadRequest,responseWeb.Code)
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := dbConf()
	truncateCategory(db)

	router := routerConf(db)

	tx,_ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Create(context.Background(),tx,domain.Category{
		Name: "Babay",
	})
	tx.Commit()

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete,URL_ENDPOINT + "/" + strconv.Itoa(category.Id),nil)
	request.Header.Add("Content-Type","applilcation/json")
	request.Header.Add("MAMAT-API-KEY","mamat-rahmat")
	router.ServeHTTP(recorder,request)

	r := recorder.Result().Body
	var responseWeb web.WebResponse
	decoder := json.NewDecoder(r)
	errDecode := decoder.Decode(&responseWeb)

	helpers.PanicIfError(errDecode)

	assert.Equal(t,http.StatusOK,responseWeb.Code)
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := dbConf()
	// truncateCategory(db)

	router := routerConf(db)

	// tx,_ := db.Begin()
	// categoryRepository := repository.NewCategoryRepository()
	// category := categoryRepository.Create(context.Background(),tx,domain.Category{
	// 	Name: "Babay",
	// })
	// tx.Commit()

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodDelete,URL_ENDPOINT + "/" + "1920",nil)
	request.Header.Add("Content-Type","applilcation/json")
	request.Header.Add("MAMAT-API-KEY","mamat-rahmat")
	router.ServeHTTP(recorder,request)

	r := recorder.Result().Body
	var responseWeb web.WebResponse
	decoder := json.NewDecoder(r)
	errDecode := decoder.Decode(&responseWeb)

	helpers.PanicIfError(errDecode)

	assert.Equal(t,http.StatusNotFound,responseWeb.Code)
}
