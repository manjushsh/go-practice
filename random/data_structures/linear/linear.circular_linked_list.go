package linear

type CircularLinkedList struct {
	head             *DLNode
	linkedListLength int
}

func (cl *CircularLinkedList) IsEmpty() bool {
	return cl.linkedListLength == 0
}

func (cl *CircularLinkedList) IsFull() bool {
	return cl.linkedListLength == LINKED_LIST_SIZE
}
