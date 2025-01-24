package linear

import "fmt"

type Queue struct {
	items       []int
	queueLength int8
}

const MAX_QUEUE_SIZE = 5

func NewQueue() *Queue {
	return &Queue{
		queueLength: 0,
	}
}

// IsEmpty checks if the queue is empty.
// It returns true if the queue has no items, otherwise false.
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Enqueue adds a new element to the end of the queue.
//
// Parameters:
//
//	numberToInsert (int): The integer value to be added to the queue.
func (q *Queue) Enqueue(numberToInsert int) {
	if len(q.items) == MAX_QUEUE_SIZE {
		fmt.Println("Queue is full.")
		// fmt.Printf("Queue is full. Replacing the first element: %d\n", q.items[0])
		// q.items = q.items[1:]
	}
	q.items = append(q.items, numberToInsert)
}

// Dequeue removes and returns the first element from the queue.
// If the queue is empty, it returns -1.
// It shifts all remaining elements one position to the left.
func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		return -1
	}
	first := q.items[0]
	q.items = q.items[1:]
	return first
}

// Peek returns the front item of the queue without removing it.
// If the queue is empty, it returns -1.
func (q *Queue) Peek() int {
	if q.IsEmpty() {
		return -1
	}
	return q.items[0]
}

func (q *Queue) Size() int {
	return len(q.items)
}

// Clear removes all elements from the queue and resets its length to zero.
func (q *Queue) Clear() {
	q.items = []int{}
	q.queueLength = 0
}

func (q *Queue) Print() {
	fmt.Println(q.items)
}
