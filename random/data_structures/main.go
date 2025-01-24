package main

import (
	"fmt"

	"github.com/manjushsh/go-practice/random/data_structures/linear"
)

func main() {
	fmt.Println("=== LINEAR START ===")
	stack := linear.NewStack()
	stack.Push(1)
	fmt.Printf("Top on stack: %d\n", stack.Peek())

	queue := linear.NewQueue()
	queue.Enqueue(6)
	queue.Enqueue(3)
	fmt.Printf("Queue Peek: %d \n", queue.Peek())
	removed := queue.Dequeue()
	fmt.Printf("Queue peek after removing %d is: %d\n", removed, queue.Peek())

	slist := linear.NewSingleLinkedList()
	slist.InsertAsTail(1)
	slist.InsertAsTail(2)
	slist.InsertAsHead(3)
	slist.InsertAfterAValue(4, 2)
	slist.InsertAtIndex(5, 2)
	fmt.Printf("List size: %d\n", slist.Size())
	fmt.Printf("List head: %d\n", slist.Peek())
	slist.RemoveByValue(2)
	slist.Print()

	fmt.Println("=== LINEAR END ===")

}
