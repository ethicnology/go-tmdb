package db

import "github.com/jmoiron/sqlx"

type Description struct {
	ID        int    `db:"id"`
	CreatedAt string `db:"created_at"`
	IdThing   int    `db:"id_thing"`
	Lang      string `db:"lang"`
	Label     string `db:"label"`
}

func AddDescription(conn *sqlx.DB, idThing int, lang string, label string) int {
	if label == "" {
		println("EMPTY: descr ", idThing, lang)
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
