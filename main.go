package main

import (
	"flag"
	"go-tmdb/api"
	"go-tmdb/collect"
)

func main() {
	mode := flag.String("mode", "collect", "api or collect")
	flag.Parse()

	if mode != nil && *mode == "collect" {
		collect.Start()
	} else {
		api.Start()
	}

}
