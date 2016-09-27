package database

import (
	"database/sql"
	"log"

	// Mysql Driver
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

//InitDB Once
func InitDB() {
	conn, err := sql.Open("mysql", "wiki:wiki@tcp(localhost:3306)/wiki?charset=utf8")
	if err != nil {
		panic("Error opening database:" + err.Error())
	}

	db = conn
	log.Printf("Database connection established")
}
