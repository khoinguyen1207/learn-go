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

	return nil
}

func ViewBooks(store *LibraryStore) {
	for _, book := range store.books {
		fmt.Printf("ID: %s | Title: %s | Author: %s | Available: %t\n", book.ID, book.Title, book.Author, !book.IsBorrowed)
	}
}

func AddBorrower(store *LibraryStore) error {
	id := utils.GenerateId()
	name := utils.GetNonEmptyString("- Enter borrower name: ")
	email := utils.GetNonEmptyString("- Enter borrower email: ")

	newBorrower := model.NewBorrower(id, name, email)
	store.borrowers[id] = *newBorrower
	fmt.Println("✅ Borrower added successfully!")

	return nil
}

func ViewBorrowers(store *LibraryStore) {
	if len(store.borrowers) == 0 {
		fmt.Println("No borrowers found.")
		return
	}
	for _, borrower := range store.borrowers {
		fmt.Printf("ID: %s | Name: %s | Email: %s\n", borrower.ID, borrower.Name, borrower.Email)
	}
}

func BorrowBook(store *LibraryStore) {
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
