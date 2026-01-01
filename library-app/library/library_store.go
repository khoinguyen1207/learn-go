package library

import "github.com/khoinguyen/learn-go/library-app/model"

type LibraryStore struct {
	books map[string]model.Book
}

func NewLibraryStore() *LibraryStore {
	return &LibraryStore{
		books: make(map[string]model.Book),
	}
}
