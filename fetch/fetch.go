package fetch

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

//hrefToURL manages local and remote href links
func hrefToURL(baseURL, href string) (string, error) {

	base, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	link, err := base.Parse(href)
	if err != nil {
		return "", err
	}

	//Cleanup the URL
	link.Fragment = ""

	if link.Scheme != "http" && link.Scheme != "https" {
		return "", errors.New("Unsupported protocol")
	}

	return link.String(), nil
}

//Fetch gets the body and the linked documents for a given URL
func Fetch(url string) ([]byte, []string, error) {
	log.Println("Fetching", url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return nil, nil, err
	}

	var urls []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if ok {
			linkurl, err := hrefToURL(url, href)
			if err != nil {
				//Skip invalid URLs
				log.Println("ERROR:", err)
				return
			}
			urls = append(urls, linkurl)
		}
	})

	return body, urls, nil
}
