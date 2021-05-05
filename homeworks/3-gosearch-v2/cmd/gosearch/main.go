package main

import (
	"flag"
	"fmt"
	"go-core-2/homeworks/3-gosearch-v2/pkg/crawler/spider"
	"go-core-2/homeworks/3-gosearch-v2/pkg/index"
	"log"
)

type search struct {
	scanner *spider.Service
	sites   []string
	depth   int
	store   *index.Store
}

func main() {
	var token = flag.String("s", "", "search for a particular word/token")
	flag.Parse()
	if *token == "" {
		fmt.Println("exiting as no token to search for was provided by input")
		return
	}

	s := new()
	fmt.Println("Processing...")

	for _, url := range s.sites {
		od, err := s.scanner.Scan(url, s.depth)
		if err != nil {
			log.Println("error when scanning a site:", err)
		}
		s.store.Append(od)
	}

	s.store.Sort()

	fmt.Println("Search results:")
	docs := s.store.Search(token)
	for _, d := range docs {
		fmt.Println(d)
	}
}

func new() *search {
	s := search{}
	s.sites = []string{"https://go.dev", "https://golang.org/"}
	s.depth = 2
	s.scanner = spider.New()
	s.store = index.New()
	return &s
}
