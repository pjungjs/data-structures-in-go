package godatastructures

import (
	"encoding/json"
	"fmt"
	"golang.org/x/exp/constraints"
	"hash/fnv"
)

type SetInterface[T constraints.Ordered] interface {
	hash(key T) int      // helper function to assign key to an index in the underlying array
	resize()             // helper function to double the size of the underlying array
	Contains(key T) bool // returns whether set contains element key. O(1)
	Insert(key T)        // inserts key in set if not present, increases size by 1. O(1)
	Remove(key T)        // removes item from set if present, decreases size by 1. O(1)
	Empty() bool         // returns whether set is empty. O(1)
	Size() int           // returns number of elements in set. O(1)
	Values() []T         // returns all values in the set.
}

// Set should implement SetInterface; write a constructor & methods to complete it
type Set[T constraints.Ordered] struct {
	arr     []*BinarySearchTree[T]
	maxFill float32 // value in (0, 1] indicating the maximum allowable ratio of elements to arr before arr is resized
	size    int
}

// /*      Methods: Set struct      */
func NewSet[T constraints.Ordered]() *Set[T] {
	arr := make([]*BinarySearchTree[T], 4)
	for i := range arr {
		arr[i] = NewBinarySearchTree[T]()
	}
	return &Set[T]{arr, 0.75, 0}
}

func SetFromSlice[T constraints.Ordered](vals []T) *Set[T] {
	arr := make([]*BinarySearchTree[T], len(vals))
	for i := range arr {
		arr[i] = NewBinarySearchTree[T]()
	}
	set := &Set[T]{arr, 0.75, 0}
	for _, val := range vals {
		set.Insert(val)
	}
	return set
}

// helper function to assign key to an index in the underlying array
func (s *Set[T]) hash(key T) int {
	//create a new 32-bit FNV-1 hash.Hash
	h := fnv.New32()

	//marshal key into JSON format's bytes
	bytes, _ := json.Marshal(key)

	//write bytes into hash
	h.Write(bytes)

	//convert uint32 hash to an int, then find the index number
	return int(h.Sum32() % uint32(cap(s.arr)))
}

// helper function to double the size of the underlying array
func (s *Set[T]) resize() {
	newLen := cap(s.arr) * 2
	if cap(s.arr) == 0 {
		newLen = 1
	}

	//save the current keys, since they need to be rehashed into the new array
	oldVals := s.Values()

	newArr := make([]*BinarySearchTree[T], newLen)
	for i := range newArr {
		newArr[i] = NewBinarySearchTree[T]()
	}
	s.arr = newArr

	//the index depends on the size of arr, so every key must be re-inserted
	for _, key := range oldVals {
		s.arr[s.hash(key)].Insert(key, nil)
	}
}

// helper function to update the maxFill ratio: keeps it in (0, 1], resets to the 0.75 default if not
func (s *Set[T]) updateMaxFill() {
	if s.maxFill <= 0 || s.maxFill > 1 {
		s.maxFill = 0.75
	}
}

// returns whether set contains element key. O(1)
func (s *Set[T]) Contains(key T) bool {
	if cap(s.arr) == 0 {
		return false
	}

	//first find the right BinarySearchTree,
	//then check if the Tree Contains the key
	bst := s.arr[s.hash(key)]
	return bst.Contains(key)
}

// inserts key in set if not present, increases size by 1. O(1)
func (s *Set[T]) Insert(key T) {
	if s.Contains(key) {
		fmt.Printf("The Key: %v already exist on the Set\n", key)
		return
	}
	s.updateMaxFill()
	if cap(s.arr) == 0 || float32(s.Size())/float32(cap(s.arr)) >= s.maxFill {
		s.resize()
	}

	bst := s.arr[s.hash(key)]
	bst.Insert(key, nil) //func (bst *BinarySearchTree[T]) Insert(key T, val any) {...}
	s.size++
}

// removes item from set if present, decreases size by 1. O(1)
func (s *Set[T]) Remove(key T) {
	if !s.Contains(key) {
		fmt.Printf("The Key: %v does not exist on the Set\n", key)
		return
	}

	bst := s.arr[s.hash(key)]
	bst.Remove(key)
	s.size--
}

// returns whether set is empty. O(1)
func (s *Set[T]) Empty() bool {
	return s.size == 0
}

// returns number of elements in set. O(1)
func (s *Set[T]) Size() int {
	return s.size
}

// returns all values in the set.
func (s *Set[T]) Values() []T {
	var allVals []T
	for _, bst := range s.arr {
		pairs := bst.InOrderTraversal()
		for _, pair := range pairs {
			allVals = append(allVals, pair.Key)
		}
	}
	return allVals
}
