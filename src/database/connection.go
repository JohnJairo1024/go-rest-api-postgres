package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:secure_pass_here@localhost/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
