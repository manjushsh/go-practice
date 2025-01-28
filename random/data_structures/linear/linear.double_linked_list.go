package linear

import "fmt"

type DLNode struct {
	data     int
	next     *DLNode
	previous *DLNode
}

type DoubleLinkedList struct {
	head             *DLNode
	linkedListLength int
}

const MAX_DOUBLE_LINKED_LIST_SIZE = 5

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{}
}

func (l *DoubleLinkedList) IsEmpty() bool {
	return l.linkedListLength == 0
}

func (l *DoubleLinkedList) IsFull() bool {
	return l.linkedListLength == MAX_DOUBLE_LINKED_LIST_SIZE
}

func (l *DoubleLinkedList) createANode(numberToInsert int) *DLNode {
	return &DLNode{
		data:     numberToInsert,
		next:     nil,
		previous: nil,
	}
}

func (l *DoubleLinkedList) Print() {
	fmt.Print("Linked List data: ")
	for currentNode := l.head; currentNode != nil; currentNode = currentNode.next {
		fmt.Printf("%d ", currentNode.data)
	}
	fmt.Println()
}

func (l *DoubleLinkedList) Insert(valueToStore int) int {
	if l.IsFull() {
		fmt.Println("Linked List is FULL.")
		return -1
	}

	newNode := l.createANode(valueToStore)
	if l.IsEmpty() {
		l.head = newNode
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
		newNode.previous = currentNode
	}
	l.linkedListLength++
	return l.linkedListLength - 1
}

func (l *DoubleLinkedList) InsertAsHead(valueToStore int) int {
	if l.IsFull() {
		fmt.Println("Linked List is FULL.")
		return -1
	}

	newNode := l.createANode(valueToStore)
	if !l.IsEmpty() {
		l.head.previous = newNode
		newNode.next = l.head
	}
	l.head = newNode
	l.linkedListLength++
	return 0
}

func (l *DoubleLinkedList) InsertBeforeAValue(valueToStore, valueToFind int) int {
	if l.IsEmpty() {
		fmt.Println("Linked List is empty.")
		return -1
	}
	if l.IsFull() {
		fmt.Println("Linked List is FULL.")
		return -1
	}

	index := 0
	for currentNode := l.head; currentNode != nil; currentNode = currentNode.next {
		if currentNode.data == valueToFind {
			newNode := l.createANode(valueToStore)
			newNode.next = currentNode
			newNode.previous = currentNode.previous
			if currentNode.previous != nil {
				currentNode.previous.next = newNode
			} else {
				l.head = newNode
			}
			currentNode.previous = newNode
			l.linkedListLength++
			return index
		}
		index++
	}
	fmt.Println("Value not found.")
	return -1
}

func (l *DoubleLinkedList) InsertAtIndex(numberToInsert, indexToInsertAt int) int {
	if l.IsFull() {
		fmt.Println("Linked list is full.")
		return -1
	}
	if indexToInsertAt < 0 || indexToInsertAt > l.linkedListLength {
		fmt.Println("Invalid index.")
		return -1
	}
	if indexToInsertAt == 0 {
		return l.InsertAsHead(numberToInsert)
	}
	if indexToInsertAt == l.linkedListLength {
		return l.Insert(numberToInsert)
	}

	newNode := l.createANode(numberToInsert)
	currentNode := l.head
	for i := 0; i < indexToInsertAt; i++ {
		currentNode = currentNode.next
	}
	newNode.next = currentNode
	newNode.previous = currentNode.previous
	if currentNode.previous != nil {
		currentNode.previous.next = newNode
	}
	currentNode.previous = newNode
	l.linkedListLength++
	return indexToInsertAt
}

func (l *DoubleLinkedList) RemoveHead() int {
	if l.IsEmpty() {
		fmt.Println("Linked List is empty.")
		return -1
	}
	if l.head.next == nil {
		l.head = nil
	} else {
		l.head = l.head.next
		l.head.previous = nil
	}
	l.linkedListLength--
	return 0
}

func (l *DoubleLinkedList) Remove() int {
	if l.IsEmpty() {
		fmt.Println("Linked List is empty.")
		return -1
	}

	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	if currentNode.previous != nil {
		currentNode.previous.next = nil
	} else {
		l.head = nil
	}
	l.linkedListLength--
	return l.linkedListLength
}

func (l *DoubleLinkedList) RemoveAtIndex(indexToRemoveAt int) int {
	if l.IsEmpty() {
		fmt.Println("Linked List is empty.")
		return -1
	}

	if indexToRemoveAt < 0 || indexToRemoveAt >= l.linkedListLength {
		fmt.Println("Please provide a valid index.")
		return -1
	}

	if indexToRemoveAt == 0 {
		return l.RemoveHead()
	}

	currentNode := l.head
	for i := 0; i < indexToRemoveAt; i++ {
		currentNode = currentNode.next
	}

	if currentNode.next != nil {
		currentNode.next.previous = currentNode.previous
	}
	if currentNode.previous != nil {
		currentNode.previous.next = currentNode.next
	} else {
		l.head = currentNode.next
	}
	l.linkedListLength--
	return indexToRemoveAt
}
