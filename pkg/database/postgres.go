package database

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func InitPostgres() (*sql.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=library_db sslmode=disable"

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
