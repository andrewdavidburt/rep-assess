package database

import (
	"database/sql"
	"reperio-backend-assessment/types"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB types.DatabaseDriver
)

func Setup() (err error) {
	db, err := sql.Open("sqlite3", "./weather.db")

	if err != nil {
		return
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	DB = types.NewDatabaseDriver(db)

	return
}
