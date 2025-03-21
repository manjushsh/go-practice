package linear

import "fmt"

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

func (cl *CircularLinkedList) GetSize() int {
	return cl.linkedListLength
}

func (cl *CircularLinkedList) InsertAtBeginning(data int) {
	if cl.IsFull() {
		fmt.Println("List is full")
		return
	}
	newNode := &DLNode{data: data} //	N1	null | 10 | null
	if cl.IsEmpty() {
		cl.head = newNode          //	HEAD ->	null | 10 | null
		cl.head.next = cl.head     //	10 | 10 | HEAD
		cl.head.previous = cl.head //	HEAD | 10 | HEAD
	} else {
		newNode.next = cl.head              //	N1 ->	null | 10 | HEAD
		newNode.previous = cl.head.previous //	N1->	HEAD.PREVIOUS | 10 | HEAD
		cl.head.previous.next = newNode     //	N-1 ->	N-2 | 10 | N1
		cl.head.previous = newNode          //	HEAD ->	N1 | 10 | N+1
		cl.head = newNode
	}
	cl.linkedListLength++
}

func (cl *CircularLinkedList) InsertAtEnd(data int) {
	if cl.IsFull() {
		fmt.Println("List is full")
		return
	}
	newNode := &DLNode{data: data}
	if cl.IsEmpty() {
		cl.InsertAtBeginning(data)
	} else {
		newNode.next = cl.head
		newNode.previous = cl.head.previous
		cl.head.previous.next = newNode
		cl.head.previous = newNode
	}
	cl.linkedListLength++
}

func (cl *CircularLinkedList) InsertAfter(data, previousData int) {
	if cl.IsFull() {
		fmt.Println("List is full")
		return
	}
	newNode := &DLNode{data: data}
	if cl.IsEmpty() {
		fmt.Println("List is empty")
		return
	}
	currentNode := cl.head
	for currentNode.data != previousData {
		currentNode = currentNode.next
		if currentNode == cl.head {
			fmt.Println("Previous data not found")
			return
		}
	}
	newNode.next = currentNode.next
	newNode.previous = currentNode
	currentNode.next.previous = newNode
	currentNode.next = newNode
	cl.linkedListLength++
}

func (cl *CircularLinkedList) InsertBefore(data, nextData int) {
	if cl.IsFull() {
		fmt.Println("List is full")
		return
	}
	newNode := &DLNode{data: data}
	if cl.IsEmpty() {
		fmt.Println("List is empty")
		return
	}
	currentNode := cl.head
	for currentNode.data != nextData {
		currentNode = currentNode.next
		if currentNode == cl.head {
			fmt.Println("Next data not found")
			return
		}
	}
	newNode.next = currentNode
	newNode.previous = currentNode.previous
	currentNode.previous.next = newNode
	currentNode.previous = newNode
	cl.linkedListLength++
}

func (cl *CircularLinkedList) InsertAtIndex(data, index int) {
	if cl.IsFull() {
		fmt.Println("List is full")
		return
	}
	if index < 0 || index > cl.linkedListLength {
		fmt.Println("Invalid index")
		return
	}
	if index == 0 {
		cl.InsertAtBeginning(data)
	} else if index == cl.linkedListLength {
		cl.InsertAtEnd(data)
	} else {
		newNode := &DLNode{data: data}
		currentNode := cl.head
		for i := 0; i < index-1; i++ {
			currentNode = currentNode.next
		}
		newNode.next = currentNode.next
		newNode.previous = currentNode
		currentNode.next.previous = newNode
		currentNode.next = newNode
		cl.linkedListLength++
	}
}

func (cl *CircularLinkedList) DeleteFromBeginning() {
	if cl.IsEmpty() {
		fmt.Println("List is empty")
		return
	}
	if cl.linkedListLength == 1 {
		cl.head = nil
	} else {
		cl.head.next.previous = cl.head.previous
		cl.head.previous.next = cl.head.next
		cl.head = cl.head.next
	}
	cl.linkedListLength--
}

func (cl *CircularLinkedList) DeleteFromEnd() {
	if cl.IsEmpty() {
		fmt.Println("List is empty")
		return
	}
	if cl.linkedListLength == 1 {
		cl.head = nil
	} else {
		cl.head.previous.previous.next = cl.head
		cl.head.previous = cl.head.previous.previous
	}
	cl.linkedListLength--
}

func (cl *CircularLinkedList) DeleteNode(data int) {
	if cl.IsEmpty() {
		fmt.Println("List is empty")
		return
	}
	currentNode := cl.head
	for currentNode.data != data {
		currentNode = currentNode.next
		if currentNode == cl.head {
			fmt.Println("Data not found")
			return
		}
	}
	if cl.linkedListLength == 1 {
		cl.head = nil
	} else {
		currentNode.next.previous = currentNode.previous
		currentNode.previous.next = currentNode.next
		if currentNode == cl.head {
			cl.head = currentNode.next
		}
	}
	cl.linkedListLength--
}

func (cl *CircularLinkedList) DeleteAtIndex(index int) {
	if cl.IsEmpty() {
		fmt.Println("List is empty")
		return
	}
	if index < 0 || index >= cl.linkedListLength {
		fmt.Println("Invalid index")
		return
	}
	if index == 0 {
		cl.DeleteFromBeginning()
	} else if index == cl.linkedListLength-1 {
		cl.DeleteFromEnd()
	} else {
		currentNode := cl.head
		for i := 0; i < index; i++ {
			currentNode = currentNode.next
		}
		currentNode.next.previous = currentNode.previous
		currentNode.previous.next = currentNode.next
		cl.linkedListLength--
	}
}

func (cl *CircularLinkedList) Display() {
	if cl.IsEmpty() {
		fmt.Println("List is empty")
		return
	}
	currentNode := cl.head
	for i := 0; i < cl.linkedListLength; i++ {
		fmt.Print(currentNode.data, " ")
		currentNode = currentNode.next
	}
	fmt.Println()
}

func (cl *CircularLinkedList) DisplayReverse() {
	if cl.IsEmpty() {
		fmt.Println("List is empty")
		return
	}
	currentNode := cl.head.previous
	for i := 0; i < cl.linkedListLength; i++ {
		fmt.Print(currentNode.data, " ")
		currentNode = currentNode.previous
	}
	fmt.Println()
}

func (cl *CircularLinkedList) GetHead() *DLNode {
	return cl.head
}

func (cl *CircularLinkedList) GetTail() *DLNode {
	return cl.head.previous
}

func (cl *CircularLinkedList) GetNode(data int) *DLNode {
	if cl.IsEmpty() {
		fmt.Println("List is empty")
		return nil
	}
	currentNode := cl.head
	for i := 0; i < cl.linkedListLength; i++ {
		if currentNode.data == data {
			return currentNode
		}
		currentNode = currentNode.next
	}
	fmt.Println("Data not found")
	return nil
}

func (cl *CircularLinkedList) GetNodeAtIndex(index int) *DLNode {
	if cl.IsEmpty() {
		fmt.Println("List is empty")
		return nil
	}
	if index < 0 || index >= cl.linkedListLength {
		fmt.Println("Invalid index")
		return nil
	}
	currentNode := cl.head
	for i := 0; i < index; i++ {
		currentNode = currentNode.next
	}
	return currentNode
}

func (cl *CircularLinkedList) Clear() {
	cl.head = nil
	cl.linkedListLength = 0
}

func (cl *CircularLinkedList) Destroy() {
	cl.head = nil
	cl = nil
}

func (cl *CircularLinkedList) SwapNodes(data1, data2 int) {
	if cl.IsEmpty() {
		fmt.Println("List is empty")
		return
	}
	if data1 == data2 {
		return
	}
	node1 := cl.GetNode(data1)
	node2 := cl.GetNode(data2)
	if node1 == nil || node2 == nil {
		return
	}
	if node1 == cl.head {
		cl.head = node2
	} else if node2 == cl.head {
		cl.head = node1
	}
	if node1.next == node2 {
		node1.next = node2.next
		node2.previous = node1.previous
		node2.next = node1
		node1.previous = node2
	} else if node2.next == node1 {
		node2.next = node1.next
		node1.previous = node2.previous
		node1.next = node2
		node2.previous = node1
	} else {
		temp := node1.next
		node1.next = node2.next
		node2.next = temp
		temp = node1.previous
		node1.previous = node2.previous
		node2.previous = temp
	}
}

func (cl *CircularLinkedList) Reverse() {
	if cl.IsEmpty() {
		fmt.Println("List is empty")
		return
	}
	currentNode := cl.head
	for i := 0; i < cl.linkedListLength; i++ {
		temp := currentNode.next
		currentNode.next = currentNode.previous
		currentNode.previous = temp
		currentNode = temp
	}
	cl.head = cl.head.previous
}

func (cl *CircularLinkedList) IsCircular() bool {
	if cl.IsEmpty() {
		fmt.Println("List is empty")
		return false
	}
	currentNode := cl.head
	for i := 0; i < cl.linkedListLength; i++ {
		if currentNode.next == cl.head {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (cl *CircularLinkedList) Split() (*CircularLinkedList, *CircularLinkedList) {
	if cl.IsEmpty() {
		fmt.Println("List is empty")
		return nil, nil
	}
	if cl.linkedListLength == 1 {
		fmt.Println("List has only one node")
		return nil, nil
	}
	cl1 := &CircularLinkedList{}
	cl2 := &CircularLinkedList{}
	mid := cl.linkedListLength / 2
	currentNode := cl.head
	for i := 0; i < mid; i++ {
		cl1.InsertAtEnd(currentNode.data)
		currentNode = currentNode.next
	}
	for i := mid; i < cl.linkedListLength; i++ {
		cl2.InsertAtEnd(currentNode.data)
		currentNode = currentNode.next
	}
	return cl1, cl2
}

func (cl *CircularLinkedList) Merge(cl2 *CircularLinkedList) {
	if cl.IsEmpty() {
		cl.head = cl2.head
		cl.linkedListLength = cl2.linkedListLength
	} else if cl2.IsEmpty() {
		return
	} else {
		tail1 := cl.head.previous
		tail2 := cl2.head.previous
		tail1.next = cl2.head
		cl2.head.previous = tail1
		tail2.next = cl.head
		cl.head.previous = tail2
		cl.linkedListLength += cl2.linkedListLength
		cl2.head = nil
		cl2.linkedListLength = 0
	}
}

func (cl *CircularLinkedList) JosephusProblem2(step int) {
	if cl.IsEmpty() {
		fmt.Println("List is empty")
		return
	}
	currentNode := cl.head
	for cl.linkedListLength > 1 {
		for i := 1; i < step; i++ {
			currentNode = currentNode.next
		}
		fmt.Print(currentNode.data, " ")
		cl.DeleteNode(currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println("\nWinner is", currentNode.data)
}
