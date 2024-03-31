package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("mysql", "root:PASSWORD/DB?charset=utf8mb4&parseTime=True&loc=Local")
    if err != nil {
        log.Fatal(err)
    }
	
}