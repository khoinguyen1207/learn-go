package main

import (
	"fmt"

	"github.com/khoinguyen/learn-go/interfaces/cat"
	"github.com/khoinguyen/learn-go/interfaces/dog"
	"github.com/khoinguyen/learn-go/interfaces/types"
)

func MakeSound(a types.Animal) {
	fmt.Println("Name:", a.GetName())
	fmt.Println("Speak:", a.Speak())
}

func MakeSoundPlus(a types.AnimalExtra) {
	fmt.Println("Name:", a.GetName())
	fmt.Println("Speak:", a.Speak())
	fmt.Println("Eat:", a.Eat())
}

// any is an alias for interface{}
func PrintValue(val interface{}) {
	switch val.(type) {
	case int:
		fmt.Println(val)
	case string:
		fmt.Println(val)
	default:
		fmt.Println("Type invalid")
	}
}

func main() {
	dog, err := dog.New("Bull")
	if err != nil {
		panic(err)
	}
	MakeSound(dog)

	PrintValue("=========================")

	cat, err := cat.New("Tom")
	if err != nil {
		panic(err)
	}

	MakeSoundPlus(cat)

	PrintValue(5)
	PrintValue(false)
}
