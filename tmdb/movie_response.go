package tmdb

type MovieResponse struct {
	Adult               *bool                  `json:"adult"`
	BackdropPath        *string                `json:"backdrop_path"`
	BelongsToCollection *any                   `json:"belongs_to_collection"`
	Budget              *int                   `json:"budget"`
	Genres              []*Genres              `json:"genres"`
	Homepage            *string                `json:"homepage"`
	ID                  *int                   `json:"id"`
	ImdbID              *string                `json:"imdb_id"`
	OriginalLanguage    *string                `json:"original_language"`
	OriginalTitle       *string                `json:"original_title"`
	Overview            *string                `json:"overview"`
	Popularity          *float64               `json:"popularity"`
	PosterPath          *string                `json:"poster_path"`
	ProductionCompanies []*ProductionCompanies `json:"production_companies"`
	ProductionCountries []*ProductionCountries `json:"production_countries"`
	ReleaseDate         *string                `json:"release_date"`
	Revenue             *int                   `json:"revenue"`
	Runtime             *int                   `json:"runtime"`
	SpokenLanguages     []*SpokenLanguages     `json:"spoken_languages"`
	Status              *string                `json:"status"`
	Tagline             *string                `json:"tagline"`
	Title               *string                `json:"title"`
	Video               *bool                  `json:"video"`
	VoteAverage         *float64               `json:"vote_average"`
	VoteCount           *int                   `json:"vote_count"`
	Success             *bool                  `json:"success"`
	StatusCode          *int                   `json:"status_code"`
	StatusMessage       *string                `json:"status_message"`
}

type Genres struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ProductionCompanies struct {
	ID            int    `json:"id"`
	LogoPath      string `json:"logo_path"`
	Name          string `json:"name"`
	OriginCountry string `json:"origin_country"`
}

type ProductionCountries struct {
	Iso31661 string `json:"iso_3166_1"`
	Name     string `json:"name"`
}

type SpokenLanguages struct {
	EnglishName string `json:"english_name"`
	Iso6391     string `json:"iso_639_1"`
	Name        string `json:"name"`
}
