package tmdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func GetMovie(tmdb int, lang string) (*MovieResponse, error) {
	tmdbKey := os.Getenv("TMDB_KEY")
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%v?api_key=%v&language=%v", tmdb, tmdbKey, lang)

	resp, err := http.Get(url)
	if err != nil {
		log.Panicln(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panicln(err)
	}

	var movie MovieResponse

	err = json.Unmarshal([]byte(body), &movie)
	if err != nil {
		return nil, errors.New("deserialization error")
	}

	if movie.Success != nil && !*movie.Success && movie.StatusCode != nil && *movie.StatusCode != 34 {
		return nil, errors.New(*movie.StatusMessage)
	}

	return &movie, nil
}
