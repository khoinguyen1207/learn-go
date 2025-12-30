package student

import "fmt"

type Student struct {
	Id     int
	Name   string
	Class  string
	Scores []float64
}

func (s Student) GetId() int {
	return s.Id
}

func (s Student) GetInfo() string {
	return fmt.Sprintf("Id: %d | Name: %s | Class: %s | Average: %2f", s.Id, s.Name, s.Class, s.CalculateAverage())
}

func (s Student) CalculateAverage() float64 {
	var totalScore float64
	for _, s := range s.Scores {
		totalScore += s
	}
	return totalScore / float64(len(s.Scores))
}
