package main

import (
	"fmt"
	boltdb "gophercises/urlshort/db"
	urlshort "gophercises/urlshort/util"
	"net/http"
)

func main() {
	db, err := boltdb.SetupDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	dbHandler, err := urlshort.DBHandler(db, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", dbHandler)

	// Example path: go-dev
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
