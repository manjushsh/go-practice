package linear

import "fmt"

const MAX_STACK_SIZE = 5

type Stack struct {
	items []int
}

// Push adds an integer element to the top of the stack.
// It appends the given integer to the slice of items in the stack.
//
// Parameters:
//
//	i - the integer to be added to the stack
func (s *Stack) Push(i int) {
	if len((s.items)) == MAX_STACK_SIZE {
		fmt.Println("Stack is full.")
	}
	s.items = append(s.items, i)
}

// Pop removes and returns the top item from the stack.
// If the stack is empty, it returns -1.
// Returns:
//
//	int: The top item from the stack, or -1 if the stack is empty.
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}

// Peek returns the top item from the stack without removing it.
// If the stack is empty, it returns -1.
// Returns:
//
//	int: The top item from the stack, or -1 if the stack is empty.
func (s *Stack) Peek() int {
	if s.IsEmpty() {
		return -1
	}
	return s.items[len(s.items)-1]
}

// IsEmpty checks if the stack is empty.
// It returns true if the stack has no items, otherwise false.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of items in the stack.
func (s *Stack) Size() int {
	return len(s.items)
}

// Clear removes all items from the stack.
func (s *Stack) Clear() {
	s.items = []int{}
}

// NewStack creates a new stack.
func NewStack() *Stack {
	return &Stack{}
}
