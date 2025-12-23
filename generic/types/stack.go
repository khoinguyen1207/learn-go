package types

type Stack[T any] struct {
	Items   []T
	Message string
}

func (s *Stack[T]) Push(item T) {
	s.Items = append(s.Items, item)
}

func (s *Stack[T]) Pop() T {
	n := len(s.Items)
	item := s.Items[n-1]
	s.Items = s.Items[:n-1]
	return item
}
