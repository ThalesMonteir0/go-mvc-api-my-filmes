package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var (
	DB_HOST     = "DB_HOST"
	DB_PORT     = "DB_PORT"
	DB_USER     = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"
	DB_NAME     = "DB_NAME"
)

func NewPostgresConnection() (*sql.DB, error) {
	host := os.Getenv(DB_HOST)
	port := os.Getenv(DB_PORT)
	user := os.Getenv(DB_USER)
	password := os.Getenv(DB_PASSWORD)
	dbname := os.Getenv(DB_NAME)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
