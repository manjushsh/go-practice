package linear

type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop removes and returns the top item from the stack.
// If the stack is empty, it returns -1.
// Returns:
//   int: The top item from the stack, or -1 if the stack is empty.
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}
