package main

import (
	"log"
	"time"

	"github.com/riviera-golang/201609-PracticalIntroduction/fetch"
)

func crawl(url string, depth int, visitedURLs fetch.URLCollection) {
	if depth <= 0 {
		return
	}

	_, urls, err := fetch.Fetch(url)
	if err != nil {
		log.Fatal(err)
	}

	for _, u := range urls {
		urlAlreadyVisited := visitedURLs.Contains(u)
		visitedURLs.Inc(u)
		if !urlAlreadyVisited {
			go crawl(u, depth-1, visitedURLs)
		}
	}
	return
}

func main() {

	visitedURLs := fetch.NewURLCollectionSafe()

	crawl("http://golang.org", 3, visitedURLs)

	time.Sleep(10 * time.Second)

	visitedURLs.Print(3)
}
