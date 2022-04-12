package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Handler struct {
	DB *sql.DB
}

func Init(url string) Handler {
	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatal(err)
	}

	return Handler{db}
}
