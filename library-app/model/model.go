package model

import "time"

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Borrower struct {
	ID    string
	Name  string
	Email string
}

type Transaction struct {
	TransactionId string
	BookID        string
	BorrowerID    string
	BorrowDate    time.Time
	ReturnDate    time.Time
}

func NewBook(id, title, author string) *Book {
	return &Book{
		ID:         id,
		Title:      title,
		Author:     author,
		IsBorrowed: false,
	}
}

func NewBorrower(id, name, email string) *Borrower {
	return &Borrower{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

func NewTransaction(transactionId, bookID, borrowerID string) *Transaction {
	return &Transaction{
		TransactionId: transactionId,
		BookID:        bookID,
		BorrowerID:    borrowerID,
		BorrowDate:    time.Now(),
		ReturnDate:    time.Now(),
	}
}
