package types

type Animal interface {
	Speak() string
	GetName() string
}

type AnimalExtra interface {
	Animal
	Eat() string
}
