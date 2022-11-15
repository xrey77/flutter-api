package config

// go get github.com/go-pg/pg/v10

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// create connection with postgres db
func Connect() *sql.DB {
	conn := DotEnvVariable("DATABASE_URL")
	db, err2 := sql.Open("postgres", conn)
	if err2 != nil {
		log.Fatal(err2)

	} else {
		log.Print("connected to Postgresql..")
	}
	return db
}
