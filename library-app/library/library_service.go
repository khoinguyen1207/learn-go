package library

import (
	"fmt"
	"strings"
	"time"

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

func BorrowBook(store *LibraryStore) error {
	transactionId := utils.GenerateId()
	bookId := utils.GetNonEmptyString("- Enter book id: ")
	borrowerId := utils.GetNonEmptyString("- Enter borrower id: ")

	book, bookExist := store.books[bookId]
	if !bookExist {
		return fmt.Errorf("book with id %s does not exist", bookId)
	}

	if book.IsBorrowed {
		return fmt.Errorf("book with id %s is already borrowed", bookId)
	}

	if _, borrowerExist := store.borrowers[borrowerId]; !borrowerExist {
		return fmt.Errorf("borrower with id %s does not exist", borrowerId)
	}

	newTx := model.NewTransaction(transactionId, bookId, borrowerId)
	store.transactions[transactionId] = *newTx
	book.IsBorrowed = true
	store.books[bookId] = book
	fmt.Println("✅ Book borrowed successfully!")

	return nil
}

func ViewBorrowHistory(store *LibraryStore) error {

	borrowerId := utils.GetNonEmptyString("- Enter borrower id: ")
	if _, borrowerExist := store.borrowers[borrowerId]; !borrowerExist {
		return fmt.Errorf("borrower with id %s does not exist", borrowerId)
	}

	var history []model.Transaction
	for _, tx := range store.transactions {
		if tx.BorrowerID == borrowerId {
			history = append(history, tx)
		}
	}

	if len(history) == 0 {
		fmt.Println("No borrow history found for this borrower.")
		return nil
	}
	for _, tx := range history {
		returnDate := "Not returned yet"
		if !tx.ReturnDate.IsZero() {
			returnDate = tx.ReturnDate.Format("2006-01-02")
		}
		fmt.Printf("Transaction ID: %s | Book ID: %s | Borrower ID: %s | Borrow Date: %s | Return Date: %s\n",
			tx.TransactionId, tx.BookID, tx.BorrowerID, tx.BorrowDate.Format("2006-01-02"), returnDate)
	}
	return nil
}

func ReturnBook(store *LibraryStore) error {
	transactionId := utils.GetNonEmptyString("- Enter transaction id: ")
	tx, txExist := store.transactions[transactionId]
	if !txExist {
		return fmt.Errorf("transaction with id %s does not exist", transactionId)
	}

	if !tx.ReturnDate.IsZero() {
		return fmt.Errorf("book for transaction id %s has already been returned", transactionId)
	}

	book := store.books[tx.BookID]
	book.IsBorrowed = false
	store.books[tx.BookID] = book

	tx.ReturnDate = time.Now()
	store.transactions[transactionId] = tx
	fmt.Println("✅ Book returned successfully!")

	return nil
}

func SearchBooks(store *LibraryStore) {
	query := utils.GetNonEmptyString("- Enter search query (title or author): ")
	query = strings.ToLower(query)

	var results []model.Book
	for _, book := range store.books {
		if strings.Contains(strings.ToLower(book.Title), query) || strings.Contains(strings.ToLower(book.Author), query) {
			results = append(results, book)
		}
	}

	if len(results) == 0 {
		fmt.Println("No books found matching the query.")
		return
	}

	for _, book := range results {
		fmt.Printf("ID: %s | Title: %s | Author: %s | Available: %t\n", book.ID, book.Title, book.Author, !book.IsBorrowed)
	}
}
