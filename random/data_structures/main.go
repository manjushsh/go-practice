package main

import (
	"fmt"

	"github.com/manjushsh/go-practice/random/data_structures/algorithms"
	"github.com/manjushsh/go-practice/random/data_structures/linear"
)

func testLinear() {
	fmt.Println("=== LINEAR START ===")
	fmt.Println(">> STACK <<")
	stack := linear.NewStack()
	stack.Push(1)
	fmt.Printf("Top on stack: %d\n", stack.Peek())

	fmt.Println(">> QUEUE <<")
	queue := linear.NewQueue()
	queue.Enqueue(6)
	queue.Enqueue(3)
	fmt.Printf("Queue Peek: %d \n", queue.Peek())
	removed := queue.Dequeue()
	fmt.Printf("Queue peek after removing %d is: %d\n", removed, queue.Peek())

	fmt.Println(">> SINGLE LINKED LIST <<")
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

	fmt.Println(">> DOUBLE LINKED LIST <<")
	dlist := linear.NewDoubleLinkedList()
	dlist.Insert(1)
	dlist.Insert(2)
	dlist.Insert(3)
	dlist.InsertAsHead(4)
	dlist.InsertBeforeAValue(5, 3)
	dlist.Print()

	fmt.Println("=== LINEAR END ===")
}

func testAlgorithms() {

	fmt.Println("=== ALGORITHMS START ===")
	list := []int{1, 2, 3, 4, 5}

	// Test IsNumInList
	fmt.Println(algorithms.IsNumInList(list, 3))
	fmt.Println(algorithms.IsNumInList(list, 6))

	// Test SumAllNumList
	fmt.Println(algorithms.SumAllNumList(list))
	fmt.Println(algorithms.SumAllNumList(nil))
	fmt.Println(algorithms.SumAllNumListRecursive(list))

	// Test ReverseAString
	fmt.Println(algorithms.ReverseAString("hello"))
	fmt.Println(algorithms.ReverseAString2("hello"))
	fmt.Println(algorithms.ReverseAString3("hello"))
	fmt.Println(algorithms.ReverseAString4("知己知彼，百战不殆；不知己不知敌，一胜一负；不知敌不知己，每战必殆。"))

	// Fizz Buzz
	fmt.Println(algorithms.FizzBuzz(15))

	// Decimal to any base
	fmt.Println(algorithms.DecimalToAnyBase(6, 2))
	fmt.Println(algorithms.DecimalToAnyBase(14, 8))
	fmt.Println(algorithms.DecimalToAnyBase(133, 16))

}

func main() {
	testAlgorithms()
}
