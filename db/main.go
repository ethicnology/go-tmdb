package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectionString() string {
	connectionString := os.Getenv("PG_CONNECT")
	return connectionString
}

func Connect() (*sqlx.DB, error) {
	db := sqlx.MustConnect("postgres", ConnectionString()+"?sslmode=disable")
	fmt.Println("Connected to the database")
	return db, nil
}
