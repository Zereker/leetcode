package main

import "github.com/Zereker/LeetCode/containers"

type MinStack struct {
	stack containers.Stack
	mins  containers.Stack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		stack: containers.NewStack(),
		mins:  containers.NewStack(),
	}
}

func (s *MinStack) Push(x int) {
	s.stack.Push(x)

	if s.mins.Size() == 0 {
		s.mins.Push(x)
		return
	}

	value, err := s.mins.Top()
	if err != nil {
		panic(err)
	}

	v := value.(int)
	if x > v {
		s.mins.Push(v)
	} else {
		s.mins.Push(x)
	}
}

func (s *MinStack) Pop() {
	if _, err := s.stack.Pop(); err != nil {
		panic(err)
	}

	if _, err := s.mins.Pop(); err != nil {
		panic(err)
	}
}

func (s *MinStack) Top() int {
	value, err := s.mins.Top()

	if err != nil {
		panic(err)
	}

	return value.(int)
}

func (s *MinStack) GetMin() int {
	top, err := s.mins.Top()
	if err != nil {
		panic(err)
	}

	return top.(int)
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	println(obj.GetMin())
	obj.Pop()
	obj.Top()
	println(obj.GetMin())
}
