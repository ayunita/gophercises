package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"gophercises/sitemap"
)

func main() {
	urlFlag := flag.String("url", "https://eleventy-base-blog.netlify.app/", "The url that you want to build the sitemap")
	maxDepth := flag.Int("depth", 3, "The maximum number of links to follow when building a sitemap")
	flag.Parse()

	s := sitemap.BFS(*urlFlag, *maxDepth)

	output, err := sitemap.XMLEncode(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(xml.Header), string(output))
}
