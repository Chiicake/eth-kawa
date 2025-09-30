package infra

import (
	"database/sql"
)

var GlobalDB *sql.DB

func InitDB() {
	// init a sqlite database as a global db
	db, err := sql.Open("sqlite3", "./eth-kawa.db")
	if err != nil {
		panic(err)
	}
	GlobalDB = db
}
