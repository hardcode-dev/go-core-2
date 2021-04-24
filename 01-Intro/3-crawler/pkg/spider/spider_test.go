// Package spider реализует сканер содержимого веб-сайтов.

// Пакет позволяет получить список ссылок и заголовков страниц внутри веб-сайта по его URL.

package spider

import (
	"fmt"
	"math"
	"testing"
)

func TestScanSite(t *testing.T) {
	const url = "https://go.dev"
	const depth = 2
	data, err := Scan(url, depth)
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range data {
		t.Logf("%s -> %s\n", k, v)
	}
}

func GetCircleSquare(radius float64) (Square float64) {
	if radius < 0 {
		fmt.Println("радиус не может быть меньше нуля!")
		return -1
	} else {
		Square = math.Pi * radius * radius
		return
	}
}
