package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Start() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.GET("/", getDoc)

	r.GET("/search/:query", searchTitles)

	r.GET("/thing/:id", getThing)

	r.GET("/thing/:id/details", getThingDetails)

	r.Run()
}
