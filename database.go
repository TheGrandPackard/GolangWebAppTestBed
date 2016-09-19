package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() {
	conn, err := sql.Open("mysql", "wiki:wiki@tcp(localhost:3306)/wiki?charset=utf8")
	checkDBError(err)

	db = conn
	log.Printf("Database connection established")
}

func checkDBError(err error) {
	if err != nil {
		panic(err)
	}
}
