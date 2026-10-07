package godatastructures

import "fmt"

type DequeInterface[T any] interface {
	Front() T        // returns first item in deque. O(1)
	Back() T         // returns last item in deque. O(1)
	PushFront(val T) // adds val to front of queue, increases size by 1. O(1)
	PushBack(val T)  // adds val to back of queue, increases size by 1. O(1)
	PopFront()       // removes item from front of queue, decreases size by 1. O(1)
	PopBack()        // removes item from back of queue, decreases size by 1. O(1)
	Empty() bool     // returns whether queue is empty. O(1)
	Size() int       // returns number of elements in queue. O(1) or O(n) depending on implementation
}

// Deque should implement DequeInterface; write a constructor & methods to complete it
type Deque[T any] struct {
	list DoublyLinkedListInterface[T]
}

/*      Methods: Deque struct      */
func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{NewDoublyLinkedList[T]()}
}

// returns first item in deque. O(1)
func (d *Deque[T]) Front() T {
	if d.list.Empty() {
		fmt.Println("Deque is empty")
		var zero T
		return zero
	}
	return d.list.Head().Data
}

// returns last item in deque. O(1)
func (d *Deque[T]) Back() T {
	if d.list.Empty() {
		fmt.Println("Deque is empty")
		var zero T
		return zero
	}
	return d.list.Tail().Data
}

// adds val to front of queue, increases size by 1. O(1)
func (d *Deque[T]) PushFront(val T) {
	d.list.InsertAtFront(val)
}

// adds val to back of queue, increases size by 1. O(1)
func (d *Deque[T]) PushBack(val T) {
	d.list.InsertAtEnd(val)
}

// removes item from front of queue, decreases size by 1. O(1)
func (d *Deque[T]) PopFront() {
	d.list.RemoveAtFront()
}

// removes item from back of queue, decreases size by 1. O(1)
func (d *Deque[T]) PopBack() {
	d.list.RemoveAtEnd()
}

// returns whether queue is empty. O(1)
func (d *Deque[T]) Empty() bool {
	return d.list.Empty()
}

// returns number of elements in queue. O(1) or O(n) depending on implementation
func (d *Deque[T]) Size() int {
	return d.list.Size()
}
