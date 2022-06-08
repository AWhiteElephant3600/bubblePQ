package dao

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
	err error
)

func init(){
	DB, err = sql.Open("postgres", "user=postgres password=66666 dbname=bubble sslmode=disable")
	if err != nil {
		panic(err)
	}
}
