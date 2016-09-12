package main

import (
	"log"

	"github.com/riviera-golang/201609-PracticalIntroduction/fetch"
)

type crawler struct {
	requests  <-chan string
	foundURLs chan<- string
}

func (c crawler) Run() {

	for url := range c.requests {

		_, urls, err := fetch.Fetch(url)
		if err != nil {
			log.Fatal(err)
		}

		for _, u := range urls {
			c.foundURLs <- u
		}
	}
}

func main() {

	url := "http://golang.org"

	visitedURLs := fetch.NewURLCollection()
	requests := make(chan string)
	responses := make(chan string)

	for i := 0; i < 20; i++ {
		c := crawler{requests, responses}
		go c.Run()
	}

	backlog := []string{url}
	for i := 0; i < 10000; i++ {

		selectedURL := ""
		var sendRequestChan chan string
		if len(backlog) > 0 {
			selectedURL = backlog[0]
			sendRequestChan = requests
		}
		select {
		case sendRequestChan <- selectedURL:
			backlog = backlog[1:]

		case resp := <-responses:
			if !visitedURLs.Contains(resp) {
				backlog = append(backlog, resp)
			}
			visitedURLs.Inc(resp)
		}
	}
	close(requests)

	visitedURLs.Print(3)
}
