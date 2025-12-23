package slice

import (
	"fmt"
	"reflect"
)

func Example() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr:", arr)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("slice:", slice)

	fmt.Println("Type:", reflect.TypeOf(arr).Kind() == reflect.Array)
	fmt.Println("Type:", reflect.TypeOf(slice).Kind() == reflect.Slice)

}
