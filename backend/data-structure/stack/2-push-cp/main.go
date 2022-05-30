package main

import (
	"errors"
	"fmt"
)

// Dari inisiasi stack dengan jumlah maksimal elemen 10, cobalah implementasikan operasi push.

var ErrStackOverflow = errors.New("stack overflow")

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
		return ErrStackOverflow
	} else {
		s.Top += 1
		s.Data = append(s.Data, Elemen)
	}
	return nil
}

func main() {
	nStack := NewStack(1)

	fmt.Println(nStack.Top, nStack.Data)
	nStack.Push(1)
	fmt.Println(nStack.Top, nStack.Data, nStack.Size)

}
