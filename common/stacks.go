package common

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Peek() T {
	if s.Size() <= 0 {
		panic("Stack empty when trying to peek")
	}

	return s.data[len(s.data)-1]
}

func (s *Stack[T]) Pop() T {
	if s.Size() <= 0 {
		panic("Stack empty when trying to peek")
	}

	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return item
}

func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func NewStack[T any]() *Stack[T] {
	s := &Stack[T]{}
	s.data = make([]T, 0)
	return s
}
