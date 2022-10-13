package app

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func NewDatabase() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/go-api-dev")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
