package db

import "github.com/jmoiron/sqlx"

type Title struct {
	ID         int    `db:"id"`
	CreatedAt  string `db:"created_at"`
	IdThing    int    `db:"id_thing"`
	IsOriginal bool   `db:"is_original"`
	Lang       string `db:"lang"`
	Label      string `db:"label"`
}

func TitlesByLabel(conn *sqlx.DB, label string) []Title {
	sql := `SELECT * FROM title WHERE "label" LIKE $1 limit 10`

	titles := []Title{}
	err := conn.Select(&titles, sql, "%"+label+"%")
	if err != nil {
		println(err.Error())
	}
	println(label)
	println(len(titles))
	return titles
}
