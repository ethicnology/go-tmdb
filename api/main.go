package api

import (
	"log"
	"net/http"
	"new_app/db"
	"strconv"

	"github.com/gin-gonic/gin"
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

	r := gin.Default()

	r.GET("/search/:query", func(c *gin.Context) {
		query := c.Param("query")
		titles := db.TitlesByLabel(conn, query)
		if len(titles) == 0 {
			c.JSON(http.StatusNoContent, gin.H{})
			return
		}
		c.JSON(http.StatusOK, titles)
	})

	r.GET("/thing/:id", func(c *gin.Context) {
		id := c.Param("id")

		idThing, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		thing, _ := db.GetThingById(conn, idThing)
		if thing == nil {
			c.JSON(http.StatusNoContent, gin.H{})
			return
		}

		c.JSON(http.StatusOK, thing)
	})

	r.GET("/thing/:id/details", func(c *gin.Context) {
		id := c.Param("id")

		idThing, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
			return
		}

		thing, _ := db.GetThingDetailsById(conn, idThing)
		if thing == nil {
			c.JSON(http.StatusNoContent, gin.H{})
			return
		}

		c.JSON(http.StatusOK, thing)
	})

	r.Run()
}
