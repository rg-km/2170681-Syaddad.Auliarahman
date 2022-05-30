package main

import "errors"

// Dari inisiasi stack dengan maksimal elemen sebanyak 10, implementasikan operasi peek.

var ErrEmptyStack = errors.New("stack overflow")

type Stack struct {
	// TODO: answer here
	Top  int
	Data []int
	Size int
}

func NewStack(size int) Stack {
	stack := Stack{
		Top:  -1,
		Data: make([]int, 0),
		Size: size,
	}
	// TODO: answer here
	return stack
}

func (s *Stack) Push(Elemen int) error {
	// TODO: answer here
	if len(s.Data) == s.Size {
		return ErrEmptyStack
	} else {
		s.Top += 1
		s.Data = append(s.Data, Elemen)
	}
	return nil
}
func (s *Stack) Peek() (int, error) {
	// TODO: answer here
	if s.Top == -1 {
		return 0, ErrEmptyStack
	} else {
		return s.Data[s.Top], nil
	}
}
