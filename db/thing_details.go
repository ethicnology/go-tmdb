package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type ThingDetails struct {
	ID          int            `db:"id"`
	Tmdb        sql.NullInt32  `db:"tmdb"`
	Imdb        sql.NullString `db:"imdb"`
	CreatedAt   sql.NullString `db:"created_at"`
	Duration    sql.NullInt32  `db:"duration"`
	Release     sql.NullString `db:"release"`
	Cover       sql.NullString `db:"cover"`
	Titles      []Title        `db:"titles"`
	Description []Description  `db:"descriptions"`
}

func GetThingDetailsById(conn *sqlx.DB, id int) (*ThingDetails, error) {
	sql := `
SELECT
    t.*,
    json_agg(DISTINCT ti.*) AS titles,
    json_agg(DISTINCT d.*) AS descriptions
FROM
    public.thing t
LEFT JOIN
    public.title ti ON t.id = ti.id_thing
LEFT JOIN
    public.description d ON t.id = d.id_thing
WHERE
    t.id = $1
GROUP BY
    t.id;`

	t := ThingDetails{}
	err := conn.Get(&t, sql, id)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
