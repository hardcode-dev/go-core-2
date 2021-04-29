package engine

import (
	"go-core-2/GoSearch/pkg/crawler"
	"go-core-2/GoSearch/pkg/index"
	"go-core-2/GoSearch/pkg/storage"
)

// Engine - поисковый движок.
// Его задача - обслуживание поисковых запросов.
// функциональность:
// - обработка поискового запроса;
// - поиск документов в индексе;
// - запрос документов из хранилища;
// - возврат посиковой выдачи.

// Service - поисковый движок.
type Service struct {
	index   index.Interface
	storage storage.Interface
}

// New - конструктор.
func New(index index.Interface, storage storage.Interface) *Service {
	s := Service{
		index:   index,
		storage: storage,
	}
	return &s
}

// Search ищет документы, соответствующие поисковому запросу.
func (s *Service) Search(query string) []crawler.Document {
	if query == "" {
		return nil
	}
	ids := s.index.Search(query)
	docs := s.storage.Docs(ids)
	return docs
}
