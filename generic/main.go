package main

import (
	"cmp"
	"fmt"

	"github.com/khoinguyen/learn-go/generic/types"
)

type ApiResponse[T any] struct {
	Data    T
	Message string
}

func printValue[T any](val T) {
	fmt.Println(val)
}

// Use only for == or !=
func isEqual[T comparable](a, b T) bool {
	return a == b
}

func max[T cmp.Ordered](a, b T) T {
	if a < b {
		return b
	}
	return a
}

func main() {
	// printValue("Nguyen")
	// printValue(3)
	// printValue(true)
	// printValue(3.213)

	// printValue(isEqual("Nguyen", "Nguyen"))
	// printValue(isEqual(2, 2.3))

	// printValue(max(2, 1))
	// printValue(max(2.33, 3.2))

	// data := ApiResponse[string]{
	// 	Data:    "test",
	// 	Message: "Login success",
	// }
	// output, _ := json.Marshal(data)
	// fmt.Println("data:", string(output))

	stack := types.Stack[string]{
		Items:   []string{},
		Message: "Not found",
	}
	stack.Push("a")
	stack.Push("b")
	fmt.Println(stack.Pop()) // b
	fmt.Println(stack)

}
