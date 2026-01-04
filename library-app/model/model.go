package model

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
