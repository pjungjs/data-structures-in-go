package godatastructures

import "fmt"

type StackInterface[T any] interface {
	Top() T      // returns top item in stack. O(1)
	Push(val T)  // adds val to top of stack, increases size by 1. O(1)
	Pop()        // removes item from top of stack, decreases size by 1. O(1)
	Empty() bool // returns whether stack is empty. O(1)
	Size() int   // returns number of elements in stack. O(1) or O(n) depending on implementation
}

// Stack should implement StackInterface; write a constructor & methods to complete it
type Stack[T any] struct {
	list SinglyLinkedListInterface[T]
}

// Remember that arrays in Go are not resizable.
// AlternateStack should implement StackInterface; write a constructor & methods to complete it
type AlternateStack[T any] struct {
	arr []T // array implementation of a stack
}

/*      Methods: Stack struct      */
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{NewSinglyLinkedList[T]()}
}

// using the head as the top of the stack, so it's O(1)
// returns top item in stack. O(1)
func (s Stack[T]) Top() T {
	if s.list.Empty() {
		fmt.Println("Stack is empty")
		var zero T
		return zero
	}
	return s.list.Head().Data
}

// adds val to top of stack, increases size by 1. O(1)
func (s Stack[T]) Push(val T) {
	s.list.InsertAtFront(val)
}

// removes item from top of stack, decreases size by 1. O(1)
func (s Stack[T]) Pop() {
	s.list.RemoveAtFront()
}

// returns whether stack is empty. O(1)
func (s Stack[T]) Empty() bool {
	return s.list.Empty()
}

// returns number of elements in stack. O(1) or O(n) depending on implementation
func (s Stack[T]) Size() int {
	return s.list.Size()
}

/*      Methods: AlternateStack struct      */
func NewAlternateStack[T any]() *AlternateStack[T] {
	return &AlternateStack[T]{[]T{}}
}

// returns top item in stack. O(1)
func (as *AlternateStack[T]) Top() T {
	if as.Empty() {
		fmt.Println("AlternateStack is empty")
		var zero T
		return zero
	}
	length := len(as.arr)
	return as.arr[length-1]
}

// adds val to top of stack, increases size by 1. O(1)
func (as *AlternateStack[T]) Push(val T) {
	as.arr = append(as.arr, val)
}

// removes item from top of stack, decreases size by 1. O(1)
func (as *AlternateStack[T]) Pop() {
	if as.Empty() {
		fmt.Println("AlternateStack is empty")
		return
	}
	length := len(as.arr)
	as.arr = as.arr[:length-1]
}

// returns whether stack is empty. O(1)
func (as *AlternateStack[T]) Empty() bool {
	return len(as.arr) == 0
}

// returns number of elements in stack. O(1) or O(n) depending on implementation
func (as *AlternateStack[T]) Size() int {
	return len(as.arr)
}
