package library

import (
	"fmt"

	"github.com/khoinguyen/learn-go/library-app/model"
	"github.com/khoinguyen/learn-go/library-app/utils"
)

func AddBook(store *LibraryStore) error {
	id := utils.GenerateId()
	title := utils.GetNonEmptyString("- Enter book title: ")
	author := utils.GetNonEmptyString("- Enter book author: ")

	newBook := model.NewBook(id, title, author)
	store.books[id] = *newBook
	fmt.Println("✅ Book added successfully!")
	fmt.Printf("Book ID: %+v\n", store.books)

	return nil
}

func ViewBooks() {
	// Implementation for viewing books
}

func AddBorrower() {
	// Implementation for adding a borrower
}

func ViewBorrowers() {
	// Implementation for viewing borrowers
}

func BorrowBook() {
	// Implementation for borrowing a book
}

func ReturnBook() {
	// Implementation for returning a book
}

func ViewBorrowHistory() {
	// Implementation for viewing borrow history
}

func SearchBooks() {
	// Implementation for searching books
}
