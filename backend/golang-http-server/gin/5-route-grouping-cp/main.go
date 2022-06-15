package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Movie struct {
	Title string
}

var movies = map[int]Movie{
	1: {
		"Spiderman",
	},
	2: {
		"Joker",
	},
	3: {
		"Escape Plan",
	},
	4: {
		"Lord of the Rings",
	},
}

var MovieListHandler = func(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{}) // TODO: replace this
	c.JSON(http.StatusOK, movies)
}

var MovieGetHandler = func(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{}) // TODO: replace this
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid ID")
		return
	}

	movie, ok := movies[idInt]
	if !ok {
		c.String(http.StatusNotFound, "data not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": movie})
}

func GetRouter() *gin.Engine {
	router := gin.Default()
	// TODO: answer here
	movie := router.Group("/movie")
	{
		movie.GET("/list", MovieListHandler)
		movie.GET("/get/:id", MovieGetHandler)
	}
	return router
}

func main() {
	router := GetRouter()

	router.Run(":8080")
}
