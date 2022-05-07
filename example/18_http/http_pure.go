package main

import (
	"fmt"
	"log"
	"net/http"
)

func moviesHandler(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	fmt.Fprintf(w, "hello %s movies", method)
}

func main() {
	http.HandleFunc("/movies", moviesHandler)

	log.Fatal(http.ListenAndServe("localhost:2565", nil))
}
