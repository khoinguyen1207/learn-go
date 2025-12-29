package main

import (
	"fmt"

	"github.com/khoinguyen/learn-go/cms/student"
	"github.com/khoinguyen/learn-go/cms/utils"
)

func main() {

	for {
		utils.ClearScreen()
		fmt.Println("Welcome to the Content Management System. 🎓")
		fmt.Println("Options:")
		fmt.Println("- 1. Student Management 📚")
		fmt.Println("- 2. Teacher Management 👨")
		fmt.Println("- 3. Exit 🚪")

		choice := utils.GetPositiveInt("=> Choose an option: ✨ ")

		switch choice {
		case 1:
			student.StudentService()
		case 2:
			fmt.Println("s")
		case 3:
			fmt.Println("Exiting the CMS. Goodbye! 👋")
			return
		default:
			fmt.Println("Invalid choice, please try again. ❌")
		}

		utils.ReadInput("Please enter to continute....")
	}

}
