package main

import (
	"flag"
	"fmt"
	cyoa "gophercises/adventure"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := flag.Int("port", 8000, "Port to start web app on")
	filename := flag.String("file", "gopher.json", "JSON file for CYOA")

	// Open json file
	file, err := os.Open(*filename)
	if err != nil {
		fmt.Println("Open file error:", err)
	}
	defer file.Close()

	// Parse json data to a map of Chapter
	story, err := cyoa.ParseJSON(file)
	if err != nil {
		fmt.Println("Parse JSON error:", err)
	}

	h := cyoa.NewHandler(story, cyoa.WithPathFn(pathFn))

	fmt.Println("Starting the server on port:", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" || path == "/story" || path == "/story/" {
		path = "/story/intro"
	}
	return path[len("/story/"):]
}
