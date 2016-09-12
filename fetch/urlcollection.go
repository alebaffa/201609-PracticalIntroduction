package fetch

import (
	"fmt"
	"sort"
)

//URLCollection allows to count visited URLs
type URLCollection interface {
	Contains(url string) bool
	Inc(url string)
	Print(min int)
}

//NewURLCollection initialize a new empty URLs collection
//This object is NOT safe for conncurrent uses
func NewURLCollection() URLCollection {
	return urlCollection(make(map[string]int))
}

//urlCollection is a container for listing visted URLs
//It is safe for conccurent use.
type urlCollection map[string]int

//Contains returns wheter an URL is part of teh collection
func (c urlCollection) Contains(url string) bool {
	_, ok := c[url]
	return ok
}

//Inc increments the count for a given URL
func (c urlCollection) Inc(url string) {

	c[url] = c[url] + 1
}

//Print dumps the collection the standard output
func (c urlCollection) Print(min int) {

	invertedMap := make(map[int][]string)
	keys := []int{}

	for url, count := range c {
		if count > min {
			if _, ok := invertedMap[count]; !ok {
				keys = append(keys, count)
			}
			invertedMap[count] = append(invertedMap[count], url)
		}
	}

	sort.Ints(keys)

	for _, k := range keys {
		urls := invertedMap[k]
		for _, u := range urls {
			fmt.Printf("%d\t%s\n", k, u)
		}
	}
}
