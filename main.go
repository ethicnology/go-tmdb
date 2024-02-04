package main

import (
	"flag"
	"new_app/api"
	"new_app/collect"
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
