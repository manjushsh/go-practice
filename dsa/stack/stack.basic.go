package stack

import "fmt"

const MAX_SIZE = 5
type Stack struct {
	items int[]
}

func (s * Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s * Stack) IsFull() bool {
	return len(s.items) == MAX_SIZE
}

func (s *Stack) Push(item int) {
	if s.IsFull() {
		fmt.Println("Oops, Stack is full!")
	}
	s.items.append(s.items, item)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		fmt.Println("Oops, Stack is empty!")
		return -1
	}
	const last int = s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}
