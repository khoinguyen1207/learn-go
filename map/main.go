package main

import "fmt"

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {

	// Case 1
	chain := map[string]int{
		"ethereum": 1,
		"bnb":      57,
	}
	fmt.Println(chain)

	// Case 2
	chainM := make(map[string]int)
	chainM["ethereum"] = 1
	chainM["bnb"] = 57

	fmt.Println(chainM)

	value, exist := chain["base"]
	if exist {
		fmt.Println("Value:", value)
	} else {
		fmt.Println("Key does not exist")
	}

	for key, value := range chain {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	userMap := map[int]User{
		1: {ID: 1, Name: "Alice", Age: 30},
		2: {ID: 2, Name: "Bob", Age: 25},
	}

	fmt.Println(userMap)

	fmt.Println("================")
	studentMap := map[string][]string{
		"Math":    {"Alice", "Bob"},
		"Science": {"Charlie", "David"},
	}

	for _, val := range studentMap {
		for _, name := range val {
			fmt.Println(name)
		}
	}
}
