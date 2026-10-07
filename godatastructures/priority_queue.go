package godatastructures

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type PriorityQueueInterface[T constraints.Ordered, V any] interface {
	Empty() bool               // returns whether queue is empty. O(1)
	Size() int                 // returns number of elements in queue. O(1)
	Front() Pair[T]            // returns first item in queue. O(1)
	Enqueue(priority T, val V) // adds val to queue with priority, increases size by 1. O(logn)
	Dequeue()                  // removes the highest priority item from queue, decreases size by 1. O(logn)
}

// PriorityQueue should implement PriorityQueueInterface; write a constructor & methods to complete it
type PriorityQueue[T constraints.Ordered, V any] struct {
	heap MaxHeap[T]
}

// /*      Methods: PriorityQueue struct      */
func NewPriorityQueue[T constraints.Ordered, V any]() *PriorityQueue[T, V] {
	return &PriorityQueue[T, V]{MaxHeap[T]{}}
}

// returns whether queue is empty. O(1)
func (pq *PriorityQueue[T, V]) Empty() bool {
	return pq.heap.Empty()
}

// returns number of elements in queue. O(1)
func (pq *PriorityQueue[T, V]) Size() int {
	return pq.heap.Size()
}

// returns first item in queue. O(1)
func (pq *PriorityQueue[T, V]) Front() Pair[T] {
	if pq.Empty() {
		fmt.Println("PriorityQueue is empty")
		return Pair[T]{}
	}
	return pq.heap.arr[0]
}

// adds val to queue with priority, increases size by 1. O(logn)
func (pq *PriorityQueue[T, V]) Enqueue(priority T, val V) {
	heapifyUp(&pq.heap.arr, Pair[T]{priority, val}, "max")
}

// removes the highest priority item from queue, decreases size by 1. O(logn)
func (pq *PriorityQueue[T, V]) Dequeue() {
	pq.heap.RemoveTop()
}
