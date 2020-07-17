package containers

import "errors"

type Stack interface {
	Push(interface{})
	Top() (interface{}, error)
	Pop() (interface{}, error)
	Empty() bool
	Size() int
}

var ErrEmpty = errors.New("empty")

type stack struct {
	data []interface{}
	size int
}

func NewStack() Stack {
	return &stack{
		data: make([]interface{}, 0),
	}
}

func (s *stack) Push(data interface{}) {
	s.data = append(s.data, data)
	s.size++
}

func (s *stack) Top() (interface{}, error) {
	size := s.Size()
	if size == 0 {
		return nil, ErrEmpty
	}

	return s.data[size-1], nil
}

func (s *stack) Pop() (interface{}, error) {
	size := s.Size()
	if size == 0 {
		return nil, ErrEmpty
	}

	data := s.data[size-1]
	s.data = s.data[0 : size-1]
	s.size--

	return data, nil
}

func (s *stack) Empty() bool {
	return s.Size() == 0
}

func (s *stack) Size() int {
	return s.size
}
