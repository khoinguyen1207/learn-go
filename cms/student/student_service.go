package student

import (
	"fmt"

	"github.com/khoinguyen/learn-go/cms/utils"
)

var students []Student

func createStudent() {
	fmt.Println("=== Create Student ===")

	var scores []float64
	var id int
	for {
		id = utils.GetPositiveInt("Id: ")
		if utils.IsIdUnique(id, students) {
			break
		}
		fmt.Println("Duplicate Id")
	}
	name := utils.GetNonEmptyString("Name: ")
	class := utils.GetNonEmptyString("Class: ")
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

	students = append(students, student)

	fmt.Println("=> Create student successfully!")
}

func updateStudent() {
	fmt.Println("=== Update Student ===")
	id := utils.GetPositiveInt("Id: ")

	for index, student := range students {
		if student.Id == id {
			fmt.Printf("====== Update student with id=%d ====== \n", id)
			student := students[index]

			name := utils.GetOptionalString(fmt.Sprintf("- Name (%s): ", student.Name), student.Name)
			class := utils.GetOptionalString(fmt.Sprintf("- Class (%s): ", student.Class), student.Class)
			newScores := make([]float64, len(student.Scores))
			for i, sc := range student.Scores {
				prompt := fmt.Sprintf("- Input score %d (%2f): ", i+1, sc)
				newScores[i] = utils.GetOptionalPositiveFloat(prompt, sc)
			}

			students[index] = Student{
				Id:     id,
				Name:   name,
				Class:  class,
				Scores: newScores,
			}

			fmt.Println("Update student successfully!")
			return
		}
	}
	fmt.Println("Student not found")
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
		fmt.Println(s.GetInfo())
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

		fmt.Println()

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

		fmt.Println()
		utils.ReadInput("Please enter to continute....")
	}
}
