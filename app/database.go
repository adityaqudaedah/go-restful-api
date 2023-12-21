package app

import (
	"database/sql"
	"time"

	"github.com/adityaqudaedah/go_restful_api/helpers"
)


func NewDB() *sql.DB {
	db,errDb := sql.Open("mysql","root@tcp(localhost:3306)/go_restful_api?parseTime=true")

	helpers.PanicIfError(errDb)
	
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}