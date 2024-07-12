package app

import (
	"database/sql"
	"fmt"
	"go-restful/config"
	"go-restful/helper"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	dbConfig := config.GetConfig().Database
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DBName))
	helper.PanicIfError(err)

	return db
}
