package main

import "fmt"

func main() {
	// arr := [5]int{1, 2, 3, 4, 5}
	// fmt.Println("arr:", arr)

	// slice := []int{1, 2, 3, 4, 5}
	// fmt.Println("slice:", slice)

	// fmt.Println("Type:", reflect.TypeOf(arr).Kind() == reflect.Array)
	// fmt.Println("Type:", reflect.TypeOf(slice).Kind() == reflect.Slice)

	// // Slice with make
	// mslice1 := make([]int, 5)

	// fmt.Println("slice:", mslice1)
	// fmt.Println("lens:", len(mslice1))
	// fmt.Println("cap:", cap(mslice1))

	// mslice2 := make([]int, 3, 10)
	// fmt.Println("slice:", mslice2)
	// fmt.Println("lens:", len(mslice2))
	// fmt.Println("cap:", cap(mslice2))

			apple1 := []int{1, 2, 3, 4}
	apple2 := []int{5, 6, 7}

	apple1 = append(apple1, apple2...)
	fmt.Println(apple1)

}
