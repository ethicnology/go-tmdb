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

func CreateTables(conn *sqlx.DB) {

	thing := `CREATE TABLE IF NOT EXISTS public.thing (
		id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
		created_at timestamptz NULL DEFAULT now(),
		tmdb int4 NOT NULL,
		imdb varchar NULL,
		duration int4 NULL,
		"release" date NULL,
		cover varchar NULL,
		CONSTRAINT thing_pk PRIMARY KEY (id),
		CONSTRAINT thing_unique UNIQUE (tmdb),
		CONSTRAINT thing_unique_1 UNIQUE (imdb)
	)`

	title := `CREATE TABLE IF NOT EXISTS public.title (
		id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
		created_at timestamptz NOT NULL DEFAULT now(),
		id_thing int4 NOT NULL,
		is_original bool NOT NULL,
		lang varchar NOT NULL,
		"label" varchar NOT NULL,
		CONSTRAINT title_pk PRIMARY KEY (id),
		CONSTRAINT title_unique_by_thing_n_label UNIQUE (id_thing, label),
		CONSTRAINT title_unique_by_thing_n_lang_n_label UNIQUE (id_thing, lang, label),
		CONSTRAINT title_thing_fk FOREIGN KEY (id_thing) REFERENCES public.thing(id)
	)`

	description := `CREATE TABLE IF NOT EXISTS public.description (
		id int4 NOT NULL GENERATED ALWAYS AS IDENTITY,
		created_at timestamptz NOT NULL DEFAULT NOW(),
		id_thing int4 NOT NULL,
		lang varchar NOT NULL,
		"label" varchar NOT NULL,
		CONSTRAINT description_pk PRIMARY KEY (id),
		CONSTRAINT description_unique_by_thing_n_lang_n_label UNIQUE (id_thing,lang,"label"),
		CONSTRAINT description_thing_fk FOREIGN KEY (id_thing) REFERENCES public.thing(id)
	)`

	conn.MustExec(thing)
	println("table thing created")

	conn.MustExec(title)
	println("table title created")

	conn.MustExec(description)
	println("table description created")
}

func AddTitle(conn *sqlx.DB, idThing int, isOriginal bool, lang string, label string) int {
	if label == "" {
		return 0
	}

	sql := `INSERT INTO title (id_thing, is_original, lang, label) VALUES ($1, $2, $3, $4) RETURNING id`

	stmt, _ := conn.Preparex(sql)

	var lastInsertId int
	stmt.QueryRowx(idThing, isOriginal, lang, label).Scan(&lastInsertId)

	if lastInsertId == 0 {
		println("FAILED: title ", idThing, lang)
	} else {
		println("INSERT: title ", idThing, lang)
	}

	return lastInsertId
}

func AddDescription(conn *sqlx.DB, idThing int, lang string, label string) int {
	if label == "" {
		return 0
	}

	sql := `INSERT INTO description (id_thing, lang, label) VALUES ($1, $2, $3) RETURNING id`

	stmt, _ := conn.Preparex(sql)

	var lastInsertId int
	stmt.QueryRowx(idThing, lang, label).Scan(&lastInsertId)

	if lastInsertId == 0 {
		println("FAILED: descr ", idThing, lang)
	} else {
		println("INSERT: descr ", idThing, lang)
	}

	return lastInsertId
}
