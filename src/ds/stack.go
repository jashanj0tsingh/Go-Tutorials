package ds

import (
	"errors"
	"fmt"
)

// A simple stack implementation using slices.

//
type value interface{}

// stack
type Stack struct {
	top    int
	values []value
	size   int
}

// construct a stack
func NewStack(size int) *Stack {

	var s Stack

	s.top = 0
	s.size = size
	s.values = make([]value, 0)

	return &s
}

// check if stack is empty
func (s *Stack) isEmpty() bool {
	if s.top == 0 {
		return true
	} else {
		return false
	}
}

// check if stack is full
func (s *Stack) isFull() bool {
	if s.top == s.size {
		return true
	} else {
		return false
	}
}

// push operation, fails if stack is full
func (s *Stack) push(val value) error {
	//
	if s.isFull() {
		fmt.Println("stack is full")
		return errors.New("cannot push into a full stack")
	} else {
		s.values = append(s.values, val)
		s.top++
		return nil
	}
}

//
func (s *Stack) pop() error {
	//
	if s.isEmpty() {
		fmt.Println("stack is empty")
		return errors.New("cannot pop an empty stack")
	} else {
		s.values = s.values[:s.top-1]
		s.top--
		return nil
	}
}

//
func (s *Stack) peek() int {
	return s.top - 1
}

//
func (s *Stack) print() {
	if s.isEmpty() {
		fmt.Println("empty stack, nothing to print")
	} else {
		for i := range s.values {
			fmt.Println(s.values[i])
		}
	}
}
