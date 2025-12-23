package array

import "fmt"

func Example() {
	var numbers [3]int = [3]int{1, 2, 3}
	fmt.Println(numbers)    // [1 2 3]
	fmt.Println(numbers[0]) // 1

	arr := [...]int{10, 20, 30}
	fmt.Println(arr)
	fmt.Println(len(arr))

	// Multidimensional array
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	for i, row := range matrix {
		for j, value := range row {
			fmt.Printf("matrix[%d][%d] = %d\n", i, j, value)
		}
	}
}
