package linear

import "fmt"

// Node for Single Linked List
type SNode struct {
	data int
	next *SNode
}

type SingleLinkedList struct {
	head             *SNode
	linkedListLength int
}

const LINKED_LIST_SIZE = 5

func (l *SingleLinkedList) IsFull() bool {
	return l.linkedListLength == LINKED_LIST_SIZE
}

func (l *SingleLinkedList) Size() int {
	return l.linkedListLength
}

func (l *SingleLinkedList) IsEmpty() bool {
	return l.head == nil
}

func (l *SingleLinkedList) createANode(numberToInsert int) *SNode {
	return &SNode{data: numberToInsert}
}

// InsertAsTail inserts a new element with the given value into the single linked list head.
// If the list is empty, it creates a new head node with the given value.
// Otherwise, it traverses to the end of the list and adds the new node there.
// It returns the value of the inserted element.
//
// Parameters:
// - numberToInsert: The value to be inserted into the list.
//
// Returns:
// - The value of the inserted element.
func (l *SingleLinkedList) InsertAsTail(numberToInsert int) {
	if l.IsFull() {
		fmt.Println("LinkedList is full. Cannot insert.")
		return
	}
	newNode := l.createANode(numberToInsert)
	if l.IsEmpty() {
		l.head = newNode
	} else {
		currentNode := l.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = newNode
	}
	l.linkedListLength++
	fmt.Printf("Inserted %d to linkedlist tail.\n", numberToInsert)
}

func (l *SingleLinkedList) InsertAtIndex(numberToInsert, indexToInsertAt int) int {
	if indexToInsertAt < 0 || indexToInsertAt > l.linkedListLength {
		fmt.Println("Index out of bounds")
		return -1
	}

	if l.IsEmpty() && indexToInsertAt > 0 {
		fmt.Println("Index out of bounds")
		return -1
	}

	if indexToInsertAt == 0 {
		l.InsertAsHead(numberToInsert)
		return 0
	}

	currentNode := l.head
	for i := 0; i < indexToInsertAt-1; i++ {
		currentNode = currentNode.next
	}

	newNode := l.createANode(numberToInsert)
	newNode.next = currentNode.next
	currentNode.next = newNode
	l.linkedListLength++
	fmt.Printf("Inserted %d at index %d in linkedlist.\n", numberToInsert, indexToInsertAt)
	return indexToInsertAt
}

func (l *SingleLinkedList) InsertAsHead(numberToInsert int) {
	newNode := l.createANode(numberToInsert)
	if !l.IsEmpty() {
		newNode.next = l.head
	}
	l.head = newNode
	l.linkedListLength++
	fmt.Printf("Inserted %d to linkedlist as head.\n", numberToInsert)
}

func (l *SingleLinkedList) InsertAfterAValue(numberToInsert, valueToFind int) {
	if l.IsEmpty() {
		l.head = l.createANode(numberToInsert)
		l.linkedListLength++
		fmt.Printf("Linked list was empty, Inserted %d to linkedlist as head.\n", numberToInsert)
		return
	}
	currentNode := l.head
	for currentNode != nil {
		if currentNode.data == valueToFind {
			newNode := &SNode{data: numberToInsert, next: currentNode.next}
			currentNode.next = newNode
			l.linkedListLength++
			fmt.Printf("Inserted %d after %d in linkedlist.\n", numberToInsert, valueToFind)
			return
		}
		currentNode = currentNode.next
	}
	fmt.Printf("Couldn't find %d in linked list.\n", valueToFind)
}

func (l *SingleLinkedList) InsertBeforeAValue(numberToInsert, valueToFind int) {
	if l.IsEmpty() {
		l.head = l.createANode(numberToInsert)
		l.linkedListLength++
		fmt.Printf("Linked list was empty, Inserted %d to linkedlist as head.\n", numberToInsert)
		return
	}
	if l.head.data == valueToFind {
		newNode := &SNode{data: numberToInsert, next: l.head}
		l.head = newNode
		l.linkedListLength++
		fmt.Printf("Inserted %d before %d in linkedlist.\n", numberToInsert, valueToFind)
		return
	}
	currentNode := l.head
	for currentNode.next != nil {
		if currentNode.next.data == valueToFind {
			newNode := &SNode{data: numberToInsert, next: currentNode.next}
			currentNode.next = newNode
			l.linkedListLength++
			fmt.Printf("Inserted %d before %d in linkedlist.\n", numberToInsert, valueToFind)
			return
		}
		currentNode = currentNode.next
	}
	fmt.Printf("Couldn't find %d in linked list.\n", valueToFind)
}

func (l *SingleLinkedList) RemoveHead() {
	if l.IsEmpty() {
		fmt.Println("LinkedList is empty. Cannot delete.")
		return
	}
	l.head = l.head.next
	l.linkedListLength--
	fmt.Println("Removed head from linkedlist.")
}

func (l *SingleLinkedList) RemoveTail() {
	if l.IsEmpty() {
		fmt.Println("LinkedList is empty. Cannot delete.")
		return
	}
	if l.head.next == nil {
		l.head = nil
	} else {
		currentNode := l.head
		for currentNode.next.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = nil
	}
	l.linkedListLength--
	fmt.Println("Removed tail from linkedlist.")
}

func (l *SingleLinkedList) RemoveAtIndex(indexToRemoveAt int) int {
	if indexToRemoveAt < 0 || indexToRemoveAt >= l.linkedListLength {
		fmt.Println("Index out of bounds")
		return -1
	}
	if l.IsEmpty() {
		fmt.Println("LinkedList is empty. Cannot delete.")
		return -1
	}
	if indexToRemoveAt == 0 {
		l.RemoveHead()
		return 0
	}
	currentNode := l.head
	for i := 0; i < indexToRemoveAt-1; i++ {
		currentNode = currentNode.next
	}
	currentNode.next = currentNode.next.next
	l.linkedListLength--
	fmt.Printf("Removed element at index %d from linkedlist.\n", indexToRemoveAt)
	return indexToRemoveAt
}

func (l *SingleLinkedList) RemoveByValue(valueToRemove int) int {
	if l.IsEmpty() {
		fmt.Println("LinkedList is empty. Cannot delete.")
		return -1
	}
	if l.head.data == valueToRemove {
		l.head = l.head.next
		l.linkedListLength--
		fmt.Printf("Removed %d from linkedlist.\n", valueToRemove)
		return 0
	}
	currentNode := l.head
	index := 0
	for currentNode.next != nil {
		if currentNode.next.data == valueToRemove {
			currentNode.next = currentNode.next.next
			l.linkedListLength--
			fmt.Printf("Removed %d from linkedlist.\n", valueToRemove)
			return index
		}
		currentNode = currentNode.next
		index++
	}
	fmt.Printf("Couldn't find %d in linked list.\n", valueToRemove)
	return -1
}

func (l *SingleLinkedList) Peek() int {
	if l.IsEmpty() {
		fmt.Println("LinkedList is empty.")
		return -1
	}
	return l.head.data
}

func (l *SingleLinkedList) Print() {
	fmt.Print("Linked List data: ")
	for currentNode := l.head; currentNode != nil; currentNode = currentNode.next {
		fmt.Printf("%d ", currentNode.data)
	}
	fmt.Println()
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{}
}

func (l *SingleLinkedList) Clear() {
	l.head = nil
	l.linkedListLength = 0
	fmt.Println("Linked list cleared.")
}

func (l *SingleLinkedList) Contains(valueToFind int) bool {
	for currentNode := l.head; currentNode != nil; currentNode = currentNode.next {
		if currentNode.data == valueToFind {
			return true
		}
	}
	return false
}

func (l *SingleLinkedList) Get(indexToGetAt int) int {
	if indexToGetAt < 0 || indexToGetAt >= l.linkedListLength {
		fmt.Println("Index out of bounds")
		return -1
	}
	currentNode := l.head
	for i := 0; i < indexToGetAt; i++ {
		currentNode = currentNode.next
	}
	return currentNode.data
}

func (l *SingleLinkedList) Set(indexToSetAt, valueToSet int) int {
	if indexToSetAt < 0 || indexToSetAt >= l.linkedListLength {
		fmt.Println("Index out of bounds")
		return -1
	}
	currentNode := l.head
	for i := 0; i < indexToSetAt; i++ {
		currentNode = currentNode.next
	}
	currentNode.data = valueToSet
	return indexToSetAt
}

func (l *SingleLinkedList) IndexOf(valueToFind int) int {
	index := 0
	for currentNode := l.head; currentNode != nil; currentNode = currentNode.next {
		if currentNode.data == valueToFind {
			return index
		}
		index++
	}
	return -1
}

func (l *SingleLinkedList) LastIndexOf(valueToFind int) int {
	index, lastIndex := 0, -1
	for currentNode := l.head; currentNode != nil; currentNode = currentNode.next {
		if currentNode.data == valueToFind {
			lastIndex = index
		}
		index++
	}
	return lastIndex
}

func (l *SingleLinkedList) RemoveAll(valueToRemove int) int {
	if l.IsEmpty() {
		fmt.Println("LinkedList is empty. Cannot delete.")
		return -1
	}
	count := 0
	for l.head != nil && l.head.data == valueToRemove {
		l.head = l.head.next
		count++
	}
	currentNode := l.head
	for currentNode != nil && currentNode.next != nil {
		if currentNode.next.data == valueToRemove {
			currentNode.next = currentNode.next.next
			count++
		} else {
			currentNode = currentNode.next
		}
	}
	l.linkedListLength -= count
	fmt.Printf("Removed %d from linkedlist %d times.\n", valueToRemove, count)
	return count
}

func (l *SingleLinkedList) RemoveFirst() {
	if l.IsEmpty() {
		fmt.Println("LinkedList is empty. Cannot delete.")
		return
	}
	l.head = l.head.next
	l.linkedListLength--
	fmt.Println("Removed first element from linkedlist.")
}

func (l *SingleLinkedList) RemoveLast() {
	if l.IsEmpty() {
		fmt.Println("LinkedList is empty. Cannot delete.")
		return
	}
	if l.head.next == nil {
		l.head = nil
	} else {
		currentNode := l.head
		for currentNode.next.next != nil {
			currentNode = currentNode.next
		}
		currentNode.next = nil
	}
	l.linkedListLength--
	fmt.Println("Removed last element from linkedlist.")
}

func (l *SingleLinkedList) RemoveAllOccurrences(valueToRemove int) int {
	if l.IsEmpty() {
		fmt.Println("LinkedList is empty. Cannot delete.")
		return -1
	}
	count := 0
	for l.head != nil && l.head.data == valueToRemove {
		l.head = l.head.next
		count++
	}
	currentNode := l.head
	for currentNode != nil && currentNode.next != nil {
		if currentNode.next.data == valueToRemove {
			currentNode.next = currentNode.next.next
			count++
		} else {
			currentNode = currentNode.next
		}
	}
	l.linkedListLength -= count
	fmt.Printf("Removed all occurrences of %d from linkedlist %d times.\n", valueToRemove, count)
	return count
}

func (l *SingleLinkedList) Reverse() {
	if l.IsEmpty() {
		fmt.Println("LinkedList is empty. Cannot reverse.")
		return
	}
	var prev, next *SNode
	current := l.head
	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}
	l.head = prev
	fmt.Println("Linkedlist reversed.")
}

func (l *SingleLinkedList) GetFirst() int {
	if l.IsEmpty() {
		fmt.Println("LinkedList is empty.")
		return -1
	}
	return l.head.data
}

func (l *SingleLinkedList) GetLast() int {
	if l.IsEmpty() {
		fmt.Println("LinkedList is empty.")
		return -1
	}
	currentNode := l.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}
	return currentNode.data
}

func (l *SingleLinkedList) GetSubList(fromIndex, toIndex int) *SingleLinkedList {
	if fromIndex < 0 || fromIndex >= l.linkedListLength || toIndex < 0 || toIndex >= l.linkedListLength || fromIndex > toIndex {
		fmt.Println("Index out of bounds")
		return nil
	}
	subList := NewSingleLinkedList()
	currentNode := l.head
	for i := 0; i < fromIndex; i++ {
		currentNode = currentNode.next
	}
	for i := fromIndex; i <= toIndex; i++ {
		subList.InsertAsTail(currentNode.data)
		currentNode = currentNode.next
	}
	return subList
}

func (l *SingleLinkedList) AddAll(valuesToAdd []int) {
	for _, value := range valuesToAdd {
		l.InsertAsTail(value)
	}
}

func (l *SingleLinkedList) AddAllAtIndex(indexToAddAt int, valuesToAdd []int) {
	if indexToAddAt < 0 || indexToAddAt > l.linkedListLength {
		fmt.Println("Index out of bounds")
		return
	}
	if indexToAddAt == l.linkedListLength {
		l.AddAll(valuesToAdd)
		return
	}
	currentNode := l.head
	for i := 0; i < indexToAddAt-1; i++ {
		currentNode = currentNode.next
	}
	for _, value := range valuesToAdd {
		newNode := &SNode{data: value, next: currentNode.next}
		currentNode.next = newNode
		l.linkedListLength++
		currentNode = currentNode.next
	}
}
