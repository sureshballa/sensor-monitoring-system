package model

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres",
		"postgres://postgres:pass123@localhost/distributed"+
			"?sslmode=disable")

	if err != nil {
		panic(err.Error())
	}
}
