package main

import (
	"flag"
	"fmt"
	"go-core-2/homeworks/03-gosearch-v1/pkg/crawler"
	"go-core-2/homeworks/03-gosearch-v1/pkg/crawler/spider"
	"log"
	"strings"
)

func main() {
	s := spider.New()
	const depth = 2
	urls := []string{"https://go.dev", "https://golang.org"}
	var docs []crawler.Document
	for _, url := range urls {
		res, err := s.Scan(url, depth)
		if err != nil {
			log.Println(err)
			continue
		}
		docs = append(docs, res...)
	}
	token := flag.String("s", "", "слово для поиска")
	flag.Parse()
	if *token != "" {
		fmt.Println("Результаты поиска:")
		for _, d := range docs {
			if strings.Contains(strings.ToLower(d.Title), strings.ToLower(*token)) {
				fmt.Println(d.URL, d.Title)
			}
		}
	}
}
