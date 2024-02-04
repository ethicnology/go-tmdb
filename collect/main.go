package collect

import (
	"log"
	"math"
	"new_app/db"
	"new_app/tmdb"

	"github.com/joho/godotenv"
)

func Start() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	db.CreateTables(conn)

	t := db.GetLastThing(conn)
	println(t.Tmdb.Int32)

	lang := "en"

	for i := int(t.Tmdb.Int32); i < math.MaxInt; i++ {

		m, err := tmdb.GetMovie(i, lang)
		if err != nil {
			log.Fatal(i)
			log.Fatal(err.Error())
		}

		if m.ID == nil {
			continue
		}

		idThing := db.AddThing(conn, m.ID, m.ImdbID, m.Runtime, m.ReleaseDate, m.PosterPath)

		if idThing == 0 {
			idThing = t.ID
		}

		if idThing != 0 {
			if m.OriginalLanguage != nil && m.OriginalTitle != nil {
				db.AddTitle(conn, int(idThing), true, *m.OriginalLanguage, *m.OriginalTitle)
			}

			if m.Title != nil {
				db.AddTitle(conn, int(idThing), false, lang, *m.Title)
			}

			if m.Overview != nil {
				db.AddDescription(conn, int(idThing), lang, *m.Overview)
			}

			fr := "fr"
			if m.OriginalLanguage != nil && *m.OriginalLanguage != fr {
				french, _ := tmdb.GetMovie(i, "fr")

				if french.Title != nil {
					db.AddTitle(conn, int(idThing), false, fr, *french.Title)
				}

				if french.Overview != nil {
					db.AddDescription(conn, int(idThing), fr, *french.Overview)
				}

			}
		}

	}

}
