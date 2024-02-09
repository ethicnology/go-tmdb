# go-tmdb

Create a postgres database
Collect `TheMovieDatabase` movies one by one
Expose a simple api 

## fill your own values in .env
```sh
cp .env.example .env
```

## scrap tmdb
```sh
go run main.go
```

## api
```sh
go run main.go --mode api
```

```json
{
		"/search/:query ":     "Find a thing by querying part of a title. eg: /search/Amélie",
		"/thing/:id ":         "Return the thing according to the provided id. eg: /thing/211",
		"/thing/:id/details ": "Return all details of the provided id thing. eg: /thing/211/details",
}
```
