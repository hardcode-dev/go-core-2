package index

import (
	"fmt"
	"go-core-2/homeworks/3-gosearch-v2/pkg/crawler"
	"sort"
	"strings"
)

type crDocs []crawler.Document

// Store preserves the documents data
type Store struct {
	counter int
	docs    crDocs
	ind     map[uint64]int
}

// New creates a new store instance
func New() *Store {
	return &Store{
		counter: 0,
		docs:    make([]crawler.Document, 0, 50),
		ind:     make(map[uint64]int),
	}
}

// Append adds document items to the store
func (s *Store) Append(docs []crawler.Document) {
	for _, d := range docs {
		s.counter++
		d.ID = s.counter
		s.docs = append(s.docs, d)
		s.index(d.ID, d.Title)
	}
}

// Print prints the contents of the store
func (s *Store) Print() {
	for _, d := range s.docs {
		fmt.Println(d)
	}
}

// Search performs a search by the token passed in
func (s *Store) Search(token *string) []string {
	res := make([]string, 0, 10)
	h := hash(strings.ToLower(*token))
	id := s.ind[h]
	d := s.binarySearch(id, 0, len(s.docs))
	res = append(res, fmt.Sprintf("%d -> %s -> %s", d.ID, d.URL, d.Title))
	return res
}

// Sort sorts the store's docs array
func (s *Store) Sort() {
	sort.Sort(s.docs)
}

func (s *Store) binarySearch(id, l, r int) crawler.Document {
	if r < l {
		return crawler.Document{}
	}

	mid := l + (r-l)/2

	if s.docs[mid].ID == id {
		return s.docs[mid]
	}

	if id <= s.docs[mid].ID {
		// go left
		return s.binarySearch(id, l, mid-1)
	} else {
		// go right
		return s.binarySearch(id, mid+1, r)
	}
}

func (s *Store) index(id int, title string) {
	var h uint64
	title = strings.TrimRight(title, "\n")
	title = strings.Replace(title, "-", "", -1)
	title = strings.Replace(title, "&", "", -1)
	
	arr := strings.Split(title, " ")
	for _, t := range arr {
		h = hash(strings.ToLower(t))
		if h > 0 {
			s.ind[h] = id
		}
	}
}

// calculates polynomial hash
func hash(text string) uint64 {
	const (
		a = 123 // base value for hash
		m = 100003 // module on which hash is calculated
	)

	hashval := uint64(0)

	for _, r := range text {
		hashval = (hashval*a + uint64(r)) % m
	}

	return hashval
}

// methods below implement the sort.Interface
func (d crDocs) Len() int           { return len(d) }
func (d crDocs) Less(i, j int) bool { return d[i].ID < d[j].ID }
func (d crDocs) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
