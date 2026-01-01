package model

type Book struct {
	ID     string
	Title  string
	Author string
}

func NewBook(id, title, author string) *Book {
	return &Book{
		ID:     id,
		Title:  title,
		Author: author,
	}
}
