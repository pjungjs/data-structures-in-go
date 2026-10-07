package godatastructures

import (
	"encoding/json"
	"fmt"
	"golang.org/x/exp/constraints"
	"hash/fnv"
)

type MapInterface[T constraints.Ordered, V any] interface {
	hash(key T) int      // helper function to assign val to an index in the underlying array
	resize()             // helper function to double the size of the underlying array
	updateMaxFill()      // helper function to update the maxFill ratio (highest allowable ratio of elements to buckets)
	Contains(key T) bool // returns whether map contains element key. O(1)
	Get(key T) any       // returns value for key in map. throws error if not present. O(1)
	Set(key T, val V)    // inserts key-val pair in map if key is not present, increases size by 1. set key's value to val if present. O(1)
	Remove(key T)        // removes key from map if present, decreases size by 1. O(1)
	Empty() bool         // returns whether map is empty. O(1)
	Size() int           // returns number of elements in map. O(1)
	Values() []any       // returns all values in the map.
	Keys() []T           // returns all keys in the map.
	Objects() []Pair[T]  // returns all key-value pairs in the map.
}

// Map should implement MapInterface; write a constructor & methods to complete it
type Map[T constraints.Ordered, V any] struct {
	arr     []*BinarySearchTree[T]
	maxFill float32 // value in (0, 1] indicating the maximum allowable ratio of elements to arr before arr is resized
	size    int
}

// /*      Methods: Map struct      */
func NewMap[T constraints.Ordered, V any]() *Map[T, V] {
	arr := make([]*BinarySearchTree[T], 4)
	for i := range arr {
		arr[i] = NewBinarySearchTree[T]()
	}
	return &Map[T, V]{arr, 0.75, 0}
}

// helper function to assign val to an index in the underlying array
func (m *Map[T, V]) hash(key T) int {
	//create a new 32-bit FNV-1 hash.Hash
	h := fnv.New32()

	//marshal key into JSON format's bytes
	bytes, _ := json.Marshal(key)

	//write bytes into hash
	h.Write(bytes)

	//convert uint32 hash to an int then find the index number
	return int(h.Sum32() % uint32(cap(m.arr)))
}

// helper function to double the size of the underlying array
func (m *Map[T, V]) resize() {
	newLen := cap(m.arr) * 2
	if cap(m.arr) == 0 {
		newLen = 1
	}

	//save the current pairs, since they need to be rehashed into the new array
	oldPairs := m.Objects()

	newArr := make([]*BinarySearchTree[T], newLen)
	for i := range newArr {
		newArr[i] = NewBinarySearchTree[T]()
	}
	m.arr = newArr

	//the index depends on the size of arr, so every pair must be re-inserted
	for _, pair := range oldPairs {
		m.arr[m.hash(pair.Key)].Insert(pair.Key, pair.Value)
	}
}

// helper function to update the maxFill ratio: keeps it in (0, 1], resets to the 0.75 default if not
func (m *Map[T, V]) updateMaxFill() {
	if m.maxFill <= 0 || m.maxFill > 1 {
		m.maxFill = 0.75
	}
}

// returns whether map contains element key. O(1)
func (m *Map[T, V]) Contains(key T) bool {
	if cap(m.arr) == 0 {
		return false
	}

	//first find the right BinarySearchTree,
	//then check if the Tree Contains the key
	bst := m.arr[m.hash(key)]
	return bst.Contains(key)
}

// returns value for key in map. throws error if not present. O(1)
func (m *Map[T, V]) Get(key T) any {
	if !m.Contains(key) {
		fmt.Printf("Key %v does not exist on the Map\n", key)
		return nil
	}
	if cap(m.arr) == 0 {
		m.resize()
	}

	bst := m.arr[m.hash(key)]
	pair := bst.GetPair(key)
	return pair.Value
}

// inserts key-val pair in map if key is not present, increases size by 1. set key's value to val if present. O(1)
func (m *Map[T, V]) Set(key T, val V) {
	if cap(m.arr) == 0 {
		m.resize()
	}

	//resize before inserting a new key once maxFill is reached
	m.updateMaxFill()
	if !m.Contains(key) && float32(m.Size())/float32(cap(m.arr)) >= m.maxFill {
		m.resize()
	}

	bst := m.arr[m.hash(key)]
	if !bst.Contains(key) {
		bst.Insert(key, val)
		m.size++
	} else {
		bst.UpdateVal(key, val)
	}
}

// removes key from map if present, decreases size by 1. O(1)
func (m *Map[T, V]) Remove(key T) {
	if !m.Contains(key) {
		fmt.Printf("The Key: %v does not exist on the Map\n", key)
		return
	}

	bst := m.arr[m.hash(key)]
	bst.Remove(key)
	m.size--
}

// returns whether map is empty. O(1)
func (m *Map[T, V]) Empty() bool {
	return m.size == 0
}

// returns number of elements in map. O(1)
func (m *Map[T, V]) Size() int {
	return m.size
}

// returns all values in the map.
func (m *Map[T, V]) Values() []any {
	var allVals []any
	for _, bst := range m.arr {
		pairs := bst.InOrderTraversal()
		for _, pair := range pairs {
			allVals = append(allVals, pair.Value)
		}
	}
	return allVals
}

// returns all keys in the map.
func (m *Map[T, V]) Keys() []T {
	var allKeys []T
	for _, bst := range m.arr {
		pairs := bst.InOrderTraversal()
		for _, pair := range pairs {
			allKeys = append(allKeys, pair.Key)
		}
	}
	return allKeys
}

// returns all key-value pairs in the map.
func (m *Map[T, V]) Objects() []Pair[T] {
	var allPairs []Pair[T]
	for _, bst := range m.arr {
		pairs := bst.InOrderTraversal()
		for _, pair := range pairs {
			allPairs = append(allPairs, pair)
		}
	}
	return allPairs
}
