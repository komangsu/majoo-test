package backeendmajootest

import (
	"database/sql"
	"time"
)

func Connection() *sql.DB {
	db, err := sql.Open("mysql", "root:mysql123@tcp(localhost:3306)/dbmajoo?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
