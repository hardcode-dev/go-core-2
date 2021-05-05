package index

import (
	"fmt"
	"go-core-2/homeworks/3-gosearch-v2/pkg/crawler"
	"strings"
)
// Store preserves the documents data
type Store struct {
	counter int
	docs    []crawler.Document
}

// New creates a new store instance
func New() *Store {
	return &Store{
		counter: 1,
		docs:    make([]crawler.Document, 0, 50),
	}
}

// Append adds document items to the store
func (s *Store) Append(docs []crawler.Document) {
	for _, d := range docs {
		s.counter++
		d.ID = s.counter
		s.docs = append(s.docs, d)
	}
}

// Print prints out the contents of the store
func (s *Store) Print() {
	for _, d := range s.docs {
		fmt.Println(d)
	}
}

// Search performs a search by the token passed in
func (s *Store) Search(token *string) []string {
	res := make([]string, 0, 10)
	for _, d := range s.docs {
		if strings.Contains(strings.ToLower(d.Title), strings.ToLower(*token)) {
			res = append(res, fmt.Sprintf("%d -> %s -> %s", d.ID, d.URL, d.Title))
		}
	}
	return res
}
