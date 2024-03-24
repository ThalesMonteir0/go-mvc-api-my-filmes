package postgresql

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func InitConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=my-movies sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	println("conectado ao banco de dados!")

	return db, err
}
