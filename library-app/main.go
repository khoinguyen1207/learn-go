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
				fmt.Printf("❌ Error adding book: %v\n", err)
			}
		case 2:
			library.ViewBooks()
		case 3:
			library.AddBorrower()
		case 4:
			library.ViewBorrowers()
		case 5:
			library.BorrowBook()
		case 6:
			library.ViewBorrowHistory()
		case 7:
			library.ReturnBook()
		case 8:
			library.SearchBooks()
		case 9:
			fmt.Println("👋 Goodbye!")
			return
		default:
			fmt.Println("❌ Invalid choice. Please try again.")
		}

		utils.ReadInput("Press Enter to continue... 🚀")
	}
}
