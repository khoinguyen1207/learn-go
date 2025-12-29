package student

import (
	"fmt"

	"github.com/khoinguyen/learn-go/cms/utils"
)

var students []Student

func createStudent() {
	fmt.Println("=== Create Student ===")

	var scores []float64
	id := utils.GetPositiveInt("Id: ")
	name := utils.ReadInput("Name: ")
	class := utils.ReadInput("Class: ")
	totalScore := utils.GetPositiveInt("Total Score: ")

	for i := 1; i <= totalScore; i++ {
		score := utils.GetPositiveFloat(fmt.Sprintf("Input score %d: ", i))
		scores = append(scores, score)
	}

	student := Student{
		Id:     id,
		Name:   name,
		Class:  class,
		Scores: scores,
	}

	fmt.Println("New student: ", student)

	students = append(students, student)

	fmt.Println("=> Create student successfully!")
}

func updateStudent() {
	fmt.Println("=== Update Student ===")
}

func deleteStudent() {
	fmt.Println("=== Delete Student ===")
}

func listStudent() {
	fmt.Println("======= List student =======")
	if len(students) <= 0 {
		fmt.Println("No data")
		return
	}
	for _, s := range students {
		fmt.Println(s)
	}
}

func searchStudent() {
	fmt.Println("=== Search Student ===")
}

func StudentService() {
	for {
		utils.ClearScreen()
		fmt.Println("=== Student Management ===")
		fmt.Println("Options:")
		fmt.Println("- 1. Create Student")
		fmt.Println("- 2. Update Student")
		fmt.Println("- 3. Delete Student")
		fmt.Println("- 4. List Students")
		fmt.Println("- 5. Search Students")
		fmt.Println("- 6. Exit 🚪")

		choice := utils.GetPositiveInt("=> Choose an option: ✨ ")

		switch choice {
		case 1:
			createStudent()
		case 2:
			updateStudent()
		case 3:
			deleteStudent()
		case 4:
			listStudent()
		case 5:
			searchStudent()
		case 6:
			fmt.Println("Goodbye! 👋")
			return
		default:
			fmt.Println("Invalid choice, please try again. ❌")
		}

		utils.ReadInput("Please enter to continute....")
	}
}
