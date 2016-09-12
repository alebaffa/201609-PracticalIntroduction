package fetch

import "sync"

//NewURLCollectionSafe initialize a new empty URLs collection
//This object is safe for conncurrent uses
func NewURLCollectionSafe() URLCollection {
	return &urlCollectionSafe{
		URLCollection: NewURLCollection(),
	}
}

//urlCollectionSafe is a container for listing visted URLs
//It is safe for conccurent use.
type urlCollectionSafe struct {
	sync.RWMutex
	URLCollection
}

//Contains returns wheter an URL is part of teh collection
func (c *urlCollectionSafe) Contains(url string) bool {
	c.RLock()
	defer c.RUnlock()

	return c.URLCollection.Contains(url)
}

//Inc increments the count for a given URL
func (c *urlCollectionSafe) Inc(url string) {
	c.Lock()
	defer c.Unlock()

	c.URLCollection.Inc(url)
}

//Print dumps the urls having at least min count to the standard output
func (c *urlCollectionSafe) Print(min int) {
	c.RLock()
	defer c.RUnlock()

	c.URLCollection.Print(min)
}
