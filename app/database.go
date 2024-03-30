package app

import (
	"database/sql"
	"go-restful/helper"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/go-db")
	helper.PanicIfError(err)

	return db
}
