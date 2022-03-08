package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/moromin/go-svelte/backend/interface/database"
)

func NewSQLHandler() database.SQLHandler {
	driverName := os.Getenv("DB_DRIVER")
	dataSourceName := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_SSLMODE"))

	conn, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return conn
}
