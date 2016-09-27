package database

import (
	"log"

	// Mysql Driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

//InitDB Once
func InitDB() {
	conn, err := sqlx.Open("mysql", "wiki:wiki@tcp(localhost:3306)/wiki?charset=utf8")
	if err != nil {
		panic("Error opening database:" + err.Error())
	}

	db = conn
	log.Printf("Database connection established")
}
