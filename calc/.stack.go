package calc

import "errors"

var StackUnderflowError = errors.New("empty stack")

type Stack struct {
	top *node
	size int
}

type node struct {
	next *node
	data float64
}

func (s *Stack) Push(n float64) *Stack {
	if s.top == nil {
		s.top = &node{nil, n}
	} else {
		s.top = &node{s.top, n}
	}
	s.size++
	return s
}

func (s *Stack) Peek(n int) ([]float64, error) {
	ns := make([]float64, 0, n)
	tmp := s.top
	for i := 0; i < n; i++ {
		if tmp == nil {
			return ns, StackUnderflowError
		}
		v := tmp.data
		tmp = tmp.next
		ns = append(ns, v)
	}
	return ns, nil
}

func (s *Stack) Pop() (float64, error) {
	if s.top == nil {
		return 0, StackUnderflowError
	}
	
	tmp := s.top.data
	s.size--
	s.top = s.top.next

	return tmp, nil
}

func (s *Stack) Pop2() (float64, float64, error) {
	_, err := s.Peek(2)
	if err != nil {
		return 0, 0, err
	}

	n, err := s.Pop()
	m, err := s.Pop()

	return n, m, nil
}

func (s *Stack) All() []float64 {
	all,err := s.Peek(s.size)
	if err != nil {
		return nil
	}
	return all
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Empty() bool {
	return s.size == 0
}
