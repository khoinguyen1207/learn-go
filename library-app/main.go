package main

import (
	"fmt"

	"github.com/khoinguyen/learn-go/library-app/library"
	"github.com/khoinguyen/learn-go/library-app/utils"
)

func main() {

	store := library.NewLibraryStore()

	for {
		utils.ClearScreen()
		fmt.Println("===== Library Management System =====")
		fmt.Println("1. Add Book 📚")
		fmt.Println("2. View Books 📖")
		fmt.Println("3. Add Borrower 🧑‍💼")
		fmt.Println("4. View Borrowers 🧾")
		fmt.Println("5. Borrow Book 📥")
		fmt.Println("6. Borrow History 📜")
		fmt.Println("7. Return Book 📤")
		fmt.Println("8. Search Books 🔍")
		fmt.Println("9. Exit ❌")

		choice := utils.GetPositiveInt("✨ Choose an option: ")

		switch choice {
		case 1:
			utils.ClearScreen()
			fmt.Println("===== Add New Book =====")
			if err := library.AddBook(store); err != nil {
				fmt.Printf("❌ Error adding book: %v", err)
			}
		case 2:
			fmt.Println("===== View Books =====")
			library.ViewBooks(store)
		case 3:
			fmt.Println("===== Add New Borrower =====")
			library.AddBorrower(store)
		case 4:
			fmt.Println("===== View Borrowers =====")
			library.ViewBorrowers(store)
		case 5:
			fmt.Println("===== Borrow Book =====")
			if err := library.BorrowBook(store); err != nil {
				fmt.Printf("❌ Error borrowing book: %v", err)
			}
		case 6:
			fmt.Println("===== Borrow History =====")
			if err := library.ViewBorrowHistory(store); err != nil {
				fmt.Printf("❌ Error viewing borrow history: %v", err)
			}
		case 7:
			fmt.Println("===== Return Book =====")
			if err := library.ReturnBook(store); err != nil {
				fmt.Printf("❌ Error returning book: %v", err)
			}
		case 8:
			fmt.Println("===== Search Books =====")
			library.SearchBooks(store)
		case 9:
			fmt.Println("👋 Goodbye!")
			return
		default:
			fmt.Println("❌ Invalid choice. Please try again.")
		}

		utils.ReadInput("\nPress Enter to continue... 🚀")
	}
}
