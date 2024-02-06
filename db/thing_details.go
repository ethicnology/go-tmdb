package db

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type ThingDetails struct {
	Thing        Thing         `json:"thing"`
	Titles       []Title       `json:"titles"`
	Descriptions []Description `json:"descriptions"`
}

func GetThingDetailsByIdJson(conn *sqlx.DB, id int) (*ThingDetails, error) {
	sql := `
SELECT
    jsonb_build_object(
        'thing', to_jsonb(t),
        'titles', COALESCE(jsonb_agg(DISTINCT to_jsonb(ti)) FILTER (WHERE ti.id IS NOT NULL), '[]'),
        'descriptions', COALESCE(jsonb_agg(DISTINCT to_jsonb(d)) FILTER (WHERE d.id IS NOT NULL), '[]')    
	) AS result
FROM
    public.thing t
LEFT JOIN
    public.title ti ON t.id = ti.id_thing
LEFT JOIN
    public.description d ON t.id = d.id_thing
WHERE
    t.id = $1
GROUP BY
    t.id;
`
	var s string
	err := conn.Get(&s, sql, id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var result ThingDetails
	err = json.Unmarshal([]byte(s), &result)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	fmt.Printf("Thing ID: %d\n", result.Thing.ID)
	fmt.Printf("Title Count: %d\n", len(result.Titles))
	fmt.Printf("Description Count: %d\n", len(result.Descriptions))

	return &result, nil
}
