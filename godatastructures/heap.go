package godatastructures

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type HeapInterface[T constraints.Ordered] interface {
	Empty() bool  // returns whether heap is empty. O(1)
	Size() int    // returns size of heap. O(1)
	Peek() T      // returns value at top of heap. O(1)
	RemoveTop()   // removes element at top of heap. O(logn)
	Insert(val T) // inserts data val, increases size by 1. O(logn)
	Sorted() []T  // returns the elements in the heap in sorted order. O(nlogn)
}

// MaxHeap should implement HeapInterface; write a constructor & methods to complete it
type MaxHeap[T constraints.Ordered] struct {
	arr []Pair[T]
}

// MinHeap should implement HeapInterface; write a constructor & methods to complete it
type MinHeap[T constraints.Ordered] struct {
	arr []Pair[T]
}

// /*      Helper Functions      */
func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func compare[T constraints.Ordered](left T, right T, heapType string) bool {
	switch heapType {
	case "max":
		return left < right
	case "min":
		return left > right
	default:
		return false
	}
}

func swap[T constraints.Ordered](arr *[]Pair[T], i int, j int) {
	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
}

func heapifyUp[T constraints.Ordered](arr *[]Pair[T], val Pair[T], heapType string) {
	*arr = append(*arr, val)
	i := len(*arr) - 1

	for i != 0 && compare((*arr)[parent(i)].Key, (*arr)[i].Key, heapType) {
		swap(arr, parent(i), i)
		i = parent(i)
	}
}

func extract[T constraints.Ordered](arr *[]Pair[T], heapType string) Pair[T] {
	extracted := (*arr)[0]
	lastInd := len(*arr) - 1

	(*arr)[0] = (*arr)[lastInd]
	*arr = (*arr)[:lastInd]

	heapifyDown(arr, 0, heapType)

	return extracted
}

func heapifyDown[T constraints.Ordered](arr *[]Pair[T], i int, heapType string) {
	lastInd := len(*arr) - 1
	l, r := left(i), right(i)
	childToCompare := 0

	// while index has at least one child
	for l <= lastInd {
		//when left child is the only child OR left child is larger/smaller
		if l == lastInd || compare((*arr)[r].Key, (*arr)[l].Key, heapType) {
			childToCompare = l
		} else {
			childToCompare = r
		}

		//compare then swap
		if !compare((*arr)[childToCompare].Key, (*arr)[i].Key, heapType) {
			swap(arr, i, childToCompare)
			i = childToCompare
			l, r = left(i), right(i)
		} else {
			return
		}
	}
}

func sortHeap[T constraints.Ordered](arr *[]Pair[T], heapType string) []T {
	arrCopy := make([]Pair[T], len(*arr))
	copy(arrCopy, *arr)

	sortedArr := make([]T, len(*arr))
	for i := range arrCopy {
		sortedArr[i] = arrCopy[0].Key
		extract(&arrCopy, heapType)
	}
	return sortedArr
}

// /*      Methods: MaxHeap struct      */
func NewMaxHeap[T constraints.Ordered]() *MaxHeap[T] {
	return &MaxHeap[T]{[]Pair[T]{}}
}

// returns whether heap is empty. O(1)
func (h *MaxHeap[T]) Empty() bool {
	return len(h.arr) == 0
}

// returns size of heap. O(1)
func (h *MaxHeap[T]) Size() int {
	return len(h.arr)
}

// returns value at top of heap. O(1)
func (h *MaxHeap[T]) Peek() T {
	if h.Empty() {
		fmt.Println("Heap array is empty")
		var zero T
		return zero
	}
	return h.arr[0].Key
}

// removes element at top of heap. O(logn)
func (h *MaxHeap[T]) RemoveTop() {
	if h.Empty() {
		fmt.Println("Heap array is empty")
		return
	}

	extracted := extract(&h.arr, "max")
	fmt.Println("Removed Top Value:", extracted)
}

// inserts data val, increases size by 1. O(logn)
func (h *MaxHeap[T]) Insert(val T) {
	heapifyUp(&h.arr, Pair[T]{val, nil}, "max")
}

// returns the elements in the heap in sorted order. O(nlogn)
func (h *MaxHeap[T]) Sorted() []T {
	return sortHeap(&h.arr, "max")
}

// /*      Methods: MinHeap struct      */
func NewMinHeap[T constraints.Ordered]() *MinHeap[T] {
	return &MinHeap[T]{[]Pair[T]{}}
}

// returns whether heap is empty. O(1)
func (h *MinHeap[T]) Empty() bool {
	return len(h.arr) == 0
}

// returns size of heap. O(1)
func (h *MinHeap[T]) Size() int {
	return len(h.arr)
}

// returns value at top of heap. O(1)
func (h *MinHeap[T]) Peek() T {
	if h.Empty() {
		fmt.Println("Heap array is empty")
		var zero T
		return zero
	}
	return h.arr[0].Key
}

// removes element at top of heap. O(logn)
func (h *MinHeap[T]) RemoveTop() {
	if h.Empty() {
		fmt.Println("Heap array is empty")
		return
	}

	extracted := extract(&h.arr, "min")
	fmt.Println("Removed Top Value:", extracted)
}

// inserts data val, increases size by 1. O(logn)
func (h *MinHeap[T]) Insert(val T) {
	heapifyUp(&h.arr, Pair[T]{val, nil}, "min")
}

// returns the elements in the heap in sorted order. O(nlogn)
func (h *MinHeap[T]) Sorted() []T {
	return sortHeap(&h.arr, "min")
}
