package storage

import "go-core-2/GoSearch/pkg/crawler"

// Храниище отсканированных документов.
// Interface определяет контракт хранилища данных.
type Interface interface {
	Docs([]int) []crawler.Document
	StoreDocs([]crawler.Document) error
}
