package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Thing struct {
	ID        int            `db:"id"`
	Tmdb      sql.NullInt32  `db:"tmdb"`
	Imdb      sql.NullString `db:"imdb"`
	CreatedAt sql.NullString `db:"created_at"`
	Duration  sql.NullInt32  `db:"duration"`
	Release   sql.NullString `db:"release"`
	Cover     sql.NullString `db:"cover"`
}

func GetLastThing(conn *sqlx.DB) Thing {
	sql := `SELECT * FROM thing ORDER BY id DESC LIMIT 1`

	thing := Thing{}
	err := conn.Get(&thing, sql)
	if err != nil {
		println(err.Error())
	}

	return thing
}

func AddThing(conn *sqlx.DB, tmdb *int, imdb *string, duration *int, release *string, cover *string) int {
	var _imdb sql.NullString
	var _duration sql.NullInt32
	var _release sql.NullString
	var _cover sql.NullString

	if imdb != nil {
		_imdb = sql.NullString{String: *imdb, Valid: true}
	}
	if duration != nil {
		_duration = sql.NullInt32{Int32: int32(*duration), Valid: true}
	}
	if release != nil {
		_release = sql.NullString{String: *release, Valid: true}
	}
	if cover != nil {
		_cover = sql.NullString{String: *cover, Valid: true}
	}

	sql := `INSERT INTO thing (tmdb, imdb, duration, release, cover) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	stmt, _ := conn.Preparex(sql)

	var lastInsertId int
	stmt.QueryRowx(tmdb, _imdb, _duration, _release, _cover).Scan(&lastInsertId)

	if lastInsertId == 0 {
		println("FAILED: thing ", *tmdb)
	} else {
		println("INSERT: thing ", *tmdb)
	}

	return lastInsertId
}

func GetThingByIdThing(conn *sqlx.DB, idThing int) *Thing {
	sql := `SELECT * FROM thing WHERE id = $1`

	thing := Thing{}
	err := conn.Get(&thing, sql, idThing)
	if err != nil {
		println(err.Error())
	}

	return &thing
}
