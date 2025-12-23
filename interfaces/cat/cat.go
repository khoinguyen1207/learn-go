package cat

import (
	"errors"
	"strings"
)

type Cat struct {
	Name string `json:"name"`
}

func New(name string) (*Cat, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return nil, errors.New("Name is required!")
	}

	if len(name) > 50 {
		return nil, errors.New("Name is too long")
	}

	return &Cat{
		Name: name,
	}, nil
}

func (c *Cat) GetName() string {
	return c.Name
}

func (c *Cat) Speak() string {
	return "Mew Mew"
}

func (c *Cat) Eat() string {
	return "Fish"
}
