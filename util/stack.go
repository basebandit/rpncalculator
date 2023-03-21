package util

import (
	"errors"
	"fmt"
	"strings"
)

type Stack struct {
	head *node
	size int
}

type node struct {
	val  any
	next *node
}

// NewStack returns a stack implemented using a linkedlist under the hood.
func NewStack() *Stack {
	return &Stack{}
}

// Size returns the count of elements currently contained in the stack.
func (s *Stack) Size() int {
	return s.size
}

// Empty checks if the stack has zero elements.
func (s *Stack) Empty() bool {
	return s.size == 0
}

// Peek returns the element at the top of the stack.
func (s *Stack) Peek() (any, error) {
	if s.Empty() {
		return nil, errors.New("stack is empty")
	}
	return s.head.val, nil
}

// Push adds a new element to the top of the stack.
func (s *Stack) Push(data any) {
	s.head = &node{val: data, next: s.head}
	s.size++
}

// Pop removes the element at the top of the stack.
func (s *Stack) Pop() (any, error) {
	if s.Empty() {
		return nil, errors.New("stack is empty")
	}
	value := s.head.val
	s.head = s.head.next
	s.size--
	return value, nil
}

// String prints string representation of the stack.
func (s *Stack) String() string {
	var sb strings.Builder
	temp := s.head
	sb.WriteString("values currently held in the stack are: ")
	for temp != nil {
		sb.WriteString(fmt.Sprintf("%v ", temp.val))
		temp = temp.next
	}
	sb.WriteString("\n")
	return sb.String()
}
