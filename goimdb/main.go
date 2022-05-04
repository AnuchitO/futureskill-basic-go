package main

import (
	"log"
	"net/http"

	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	if id == "" {
		return c.JSON(http.StatusBadRequest, "id is required.")
	}

	for _, m := range movies {
		if m.ImdbID == id {
			return c.JSON(http.StatusOK, m)
		}
	}

	return c.JSON(http.StatusNotFound, id+" not found.")
}

func createParcel(c echo.Context) error {
	mv := &Movie{}
	if err := c.Bind(mv); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	for _, m := range movies {
		if mv.ImdbID == m.ImdbID {
			return echo.NewHTTPError(http.StatusBadRequest, "already exists.")
		}
	}

	movies = append(movies, *mv)

	return c.JSON(http.StatusCreated, mv)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/movies", getAllMoviesHandler)

	e.GET("/movies/:id", getMoviesByIdHandler)

	e.POST("/movies", createParcel)

	port := "2565"
	log.Println("starting... port:", port)
	log.Fatal(e.Start(":" + port))
}
