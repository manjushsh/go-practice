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
	for currentNode := l.head; currentNode != nil; {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.next
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
	currentNode := l.head
	for currentNode != nil {
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
		currentNode = currentNode.next
		index++
	}
	fmt.Println("Value not found.")
	return -1
}
