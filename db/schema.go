package db

import "github.com/jmoiron/sqlx"

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
