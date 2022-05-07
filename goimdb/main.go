package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Movie struct {
	ImdbID      string  `json:"imdbID"`
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Rating      float32 `json:"rating"`
	IsSuperHero bool    `json:"isSuperHero"`
}

var movies = []Movie{
	{
		ImdbID:      "tt4154796",
		Title:       "Avengers: Endgame",
		Year:        2019,
		Rating:      8.4,
		IsSuperHero: true,
	},
}

func getAllMoviesHandler(c echo.Context) error {
	y := c.QueryParam("year")

	if y == "" {
		return c.JSON(http.StatusOK, movies)
	}

	year, err := strconv.Atoi(y)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ms := []Movie{}
	for _, m := range movies {
		if m.Year == year {
			ms = append(ms, m)
		}
	}

	return c.JSON(http.StatusOK, ms)
}

func getMoviesByIdHandler(c echo.Context) error {
	id := c.Param("id")

	for _, m := range movies {
		if m.ImdbID == id {
			return c.JSON(http.StatusOK, m)
		}
	}

	return c.JSON(http.StatusNotFound, map[string]string{"message!": "not found"})
}

func createMoviesHandler(c echo.Context) error {
	m := &Movie{}

	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	movies = append(movies, *m)
	return c.JSON(http.StatusCreated, m)
}

func main() {
	e := echo.New()

	e.GET("/movies", getAllMoviesHandler)
	e.GET("/movies/:id", getMoviesByIdHandler)

	e.POST("/movies", createMoviesHandler)

	port := "2565"
	log.Println("starting... port:", port)

	log.Fatal(e.Start(":" + port))
}
