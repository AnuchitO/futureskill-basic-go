package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Movie struct {
	ImdbID      string  `json:"imdbID"`
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Rating      float32 `json:"rating"`
	IsSuperHero bool    `json:"isSuperHero"`
}

var movies []Movie

func moviesHandler(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	if method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "error : %v", err)
			return
		}

		t := Movie{}
		err = json.Unmarshal(body, &t)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "error: %s", err)
			return
		}

		movies = append(movies, t)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "hello %s created movies", "POST")
		return
	}

	b, err := json.Marshal(movies)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

func main() {
	http.HandleFunc("/movies", moviesHandler)

	err := http.ListenAndServe("localhost:2565", nil)
	log.Fatal(err)
}
