package library

import (
	"github.com/khoinguyen/learn-go/library-app/model"
)

type LibraryStore struct {
	books        map[string]model.Book
	borrowers    map[string]model.Borrower
	transactions map[string]model.Transaction
}

func NewLibraryStore() *LibraryStore {
	initialBooks := map[string]model.Book{
		"1": *model.NewBook("1", "The Great Gatsby", "F. Scott Fitzgerald"),
		"2": *model.NewBook("2", "To Kill a Mockingbird", "Harper Lee"),
		"3": *model.NewBook("3", "1984", "George Orwell"),
	}

	initialBorrowers := map[string]model.Borrower{
		"1": *model.NewBorrower("1", "Alice Johnson", "alice.johnson@example.com"),
		"2": *model.NewBorrower("2", "Bob Smith", "bob.smith@example.com"),
		"3": *model.NewBorrower("3", "Charlie Brown", "charlie.brown@example.com"),
	}

	return &LibraryStore{
		books:        initialBooks,
		borrowers:    initialBorrowers,
		transactions: make(map[string]model.Transaction),
	}
}
