package main

import (
	"flag"
	"fmt"
	"go-core-2/homeworks/2-gosearch-v1/pkg/crawler"
	"go-core-2/homeworks/2-gosearch-v1/pkg/crawler/spider"
	"log"
	"strings"
)

type search struct {
	scanner crawler.Interface
	sites   []string
	depth   int
	// store   storage
}

func main() {
	var token = flag.String("s", "", "search for a particular word/token")
	flag.Parse()
	if *token == "" {
		fmt.Println("exiting as no token to search for was provided by input")
		return
	}

	s := new()
	var docs []crawler.Document

	for _, url := range s.sites {
		od, err := s.scanner.Scan(url, s.depth)
		if err != nil {
			log.Println("error when scanning a site:", err)
		}
		docs = append(docs, od...)
	}

	fmt.Println("Search results:")
	for _, d := range docs {
		if strings.Contains(strings.ToLower(d.Title), strings.ToLower(*token)) {
			fmt.Println(d.URL, "->", d.Title)
		}
	}
}

func new() *search {
	s := search{}
	s.sites = []string{"https://go.dev", "https://golang.org/"}
	s.depth = 2
	s.scanner = spider.New()
	return &s
}
