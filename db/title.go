package db

import (
	"github.com/jmoiron/sqlx"
)

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

	return titles
}

func AddTitle(conn *sqlx.DB, idThing int, isOriginal bool, lang string, label string) int {
	if label == "" {
		println("EMPTY: title ", idThing, lang)
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
