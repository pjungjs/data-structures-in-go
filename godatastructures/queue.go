package godatastructures

import "fmt"

type QueueInterface[T any] interface {
	Front() T      // returns first item in queue. O(1)
	Enqueue(val T) // adds val to end of queue, increases size by 1. O(1) or O(n) depending on implementation
	Dequeue()      // removes item from front of queue, decreases size by 1. O(1)
	Empty() bool   // returns whether queue is empty. O(1)
	Size() int     // returns number of elements in queue. O(1) or O(n) depending on implementation
}

// Queue should implement QueueInterface; write a constructor & methods to complete it
type Queue[T any] struct {
	list SinglyLinkedListInterface[T]
}

type AlternateQueue[T any] struct {
	arr []T // array implementation of a queue
}

/*      Methods: Queue struct      */
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{NewSinglyLinkedList[T]()}
}

// returns first item in queue. O(1)
func (q Queue[T]) Front() T {
	if q.list.Empty() {
		fmt.Println("Queue is empty")
		var zero T
		return zero
	}
	return q.list.Head().Data
}

// adds val to end of queue, increases size by 1. O(n), since SinglyLinkedList.InsertAtEnd walks to the tail
func (q Queue[T]) Enqueue(val T) {
	q.list.InsertAtEnd(val)
}

// removes item from front of queue, decreases size by 1. O(1)
func (q Queue[T]) Dequeue() {
	q.list.RemoveAtFront()
}

// returns whether queue is empty. O(1)
func (q Queue[T]) Empty() bool {
	return q.list.Empty()
}

// returns number of elements in queue. O(1) or O(n) depending on implementation
func (q Queue[T]) Size() int {
	return q.list.Size()
}

/*      Methods: AlternateQueue struct      */
func NewAlternateQueue[T any]() *AlternateQueue[T] {
	return &AlternateQueue[T]{[]T{}}
}

// returns first item in queue. O(1)
func (aq *AlternateQueue[T]) Front() T {
	if aq.Empty() {
		fmt.Println("AlternateQueue is empty")
		var zero T
		return zero
	}
	return aq.arr[0]
}

// adds val to end of queue, increases size by 1. O(1)
func (aq *AlternateQueue[T]) Enqueue(val T) {
	aq.arr = append(aq.arr, val)
}

// removes item from front of queue, decreases size by 1. O(1)
func (aq *AlternateQueue[T]) Dequeue() {
	if aq.Empty() {
		fmt.Println("AlternateQueue is empty")
		return
	}
	aq.arr = aq.arr[1:]
}

// returns whether queue is empty. O(1)
func (aq *AlternateQueue[T]) Empty() bool {
	return len(aq.arr) == 0
}

// returns number of elements in queue. O(1) or O(n) depending on implementation
func (aq *AlternateQueue[T]) Size() int {
	return len(aq.arr)
}
