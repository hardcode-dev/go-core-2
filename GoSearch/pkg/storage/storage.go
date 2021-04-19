package storage

// Хранилище отсканированных документов.

import (
	"go-core-2/gosearch/pkg/crawler"
)

// Interface определяет контракт хранилища данных.
type Interface interface {
	Docs([]int) []crawler.Document
	StoreDocs([]crawler.Document) error
}
