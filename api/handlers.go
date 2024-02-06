package api

import (
	"net/http"
	"new_app/db"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getDoc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"/search/:query ":     "Find a thing by querying part of a title. eg: /search/Amélie",
		"/thing/:id ":         "Return the thing according to the provided id. eg: /thing/211",
		"/thing/:id/details ": "Return all details of the provided id thing. eg: /thing/211/details",
	})
}

func searchTitles(c *gin.Context) {
	conn := db.Connect()
	query := c.Param("query")
	titles := db.TitlesByLabel(conn, query)
	if len(titles) == 0 {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	c.JSON(http.StatusOK, titles)
}

func getThing(c *gin.Context) {
	conn := db.Connect()
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
}

func getThingDetails(c *gin.Context) {
	conn := db.Connect()
	id := c.Param("id")

	idThing, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	result, _ := db.GetThingDetailsByIdJson(conn, idThing)
	if result == nil {
		c.JSON(http.StatusNoContent, gin.H{})
		return
	}
	c.JSON(http.StatusOK, result)
}
