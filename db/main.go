package db

import (
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectionString() string {
	connectionString := os.Getenv("PG_CONNECT")
	return connectionString
}

func Connect() *sqlx.DB {
	db := sqlx.MustConnect("postgres", ConnectionString()+"?sslmode=disable")
	return db
}
