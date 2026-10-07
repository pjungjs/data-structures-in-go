package main

import (
	"fmt"
	ds "github.com/pjungjs/data-structures-in-go/godatastructures"
)

func main() {
	fmt.Println("\nData Structures in Go:")

	fmt.Println("\n\n\t>>> Singly Linked List <<<")
	sll := ds.NewSinglyLinkedList[int]()
	sll.InsertAtFront(9)
	snode := sll.Head()
	sll.InsertAtFront(7)
	sll.InsertAtEnd(3)
	sll.InsertAfter(4, sll.Head())
	fmt.Println(sll.Size(), "nodes") // expect: 4 nodes
	sll.Print()                      // expect: 7 -> 4 -> 9 -> 3 -> <nil>
	snode = sll.Head()
	sll.RemoveAfter(snode)           // expect: removed 4
	sll.RemoveAtFront()              // expect: removed 7
	sll.RemoveAtEnd()                // expect: removed 3
	fmt.Println(sll.Size(), "nodes") // expect: 1 nodes
	sll.Print()                      // expect: 9 -> <nil>

	fmt.Println("\n\n\n\t>>> Modified Singly Linked List <<<")
	msll := ds.NewModifiedSinglyLinkedList[int]()
	msll.InsertAtFront(9)
	msll.InsertAtFront(7)
	msll.InsertAtEnd(3)
	msll.InsertAfter(4, msll.Head())
	fmt.Println(msll.Size(), "nodes")      // expect: 4 nodes
	msll.Print()                           // expect: 7 -> 4 -> 9 -> 3 -> <nil>
	fmt.Println("Tail:", msll.Tail().Data) // expect: Tail: 3
	msll.RemoveAfter(msll.Head())          // expect: removed 4
	msll.RemoveAtFront()                   // expect: removed 7
	msll.RemoveAtEnd()                     // expect: removed 3
	fmt.Println(msll.Size(), "nodes")      // expect: 1 nodes
	msll.Print()                           // expect: 9 -> <nil>
	fmt.Println("Tail:", msll.Tail().Data) // expect: Tail: 9

	fmt.Println("\n\n\n\t>>> Doubly Linked List <<<")
	dll := ds.NewDoublyLinkedList[int]()
	dll.InsertAtFront(9)
	dnode := dll.Head()
	dll.InsertAtFront(7)
	dll.InsertAtEnd(3)
	dll.InsertAfter(4, dnode)
	dll.InsertBefore(5, dnode)
	fmt.Println(dll.Size(), "nodes") // expect: 5 nodes
	dll.Print()                      // expect: 7 -> 5 -> 9 -> 4 -> 3 -> <nil>
	dll.Remove(dnode)                // removed 9
	dll.RemoveAtFront()              // removed 7
	dll.RemoveAtEnd()                // removed 3
	fmt.Println(dll.Size(), "nodes") // expect: 2 nodes
	dll.Print()                      // expect: 5 -> 4 -> <nil>

	fmt.Println("\n\n\n\t>>> Stack <<<")
	s := ds.NewStack[int]()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	fmt.Println("Size:", s.Size()) // expect: Size 3
	for !s.Empty() {
		fmt.Print(s.Top(), ", ")
		s.Pop()
	}
	// expect: 3, 2, 1
	fmt.Println("Size:", s.Size()) // expect: Size 0

	fmt.Println("\n\n\n\t>>> Alternate Stack <<<")
	ss := ds.NewAlternateStack[int]()
	ss.Push(1)
	ss.Push(2)
	ss.Push(3)
	fmt.Println("Size:", ss.Size()) // expect: Size 3
	for !ss.Empty() {
		fmt.Print(ss.Top(), ", ")
		ss.Pop()
	}
	// expect: 3, 2, 1
	fmt.Println("\nSize", ss.Size()) // expect: Size 0

	fmt.Println("\n\n\n\t>>> Queue <<<")
	q := ds.NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Println("Size:", q.Size()) // expect: Size 3
	for !q.Empty() {
		fmt.Print(q.Front(), ", ")
		q.Dequeue()
	}
	// expect: 1, 2, 3
	fmt.Println("Size:", q.Size()) // expect: Size 0

	fmt.Println("\n\n\n\t>>> Alternate Queue <<<")
	qq := ds.NewAlternateQueue[int]()
	qq.Enqueue(1)
	qq.Enqueue(2)
	qq.Enqueue(3)
	fmt.Println("Size:", qq.Size()) // expect: Size 3
	for !qq.Empty() {
		fmt.Print(qq.Front(), ", ")
		qq.Dequeue()
	}
	// expect: 1, 2, 3
	fmt.Println("\nSize", qq.Size()) // expect: Size 0

	fmt.Println("\n\n\n\t>>> Deque <<<")
	dq := ds.NewDeque[int]()
	dq.PushBack(3)
	dq.PushFront(6)
	dq.PushFront(1)
	dq.PushBack(9)
	dq.PushFront(2)
	dq.PushBack(8)
	fmt.Println("Size:", dq.Size()) // expect: Size 6
	for !dq.Empty() {
		fmt.Print("Front:", dq.Front(), ", ")
		fmt.Println("Back:", dq.Back())
		dq.PopFront()
		dq.PopBack()
	}
	// expect: 2, 8, 1, 9, 6, 3
	fmt.Println("Size:", dq.Size()) // expect: Size 0

	fmt.Println("\n\n\n\t>>> Binary Search Tree <<<")
	bst := ds.NewBinarySearchTree[int]()
	bst.Insert(6, nil)
	bst.Insert(3, nil)
	bst.Insert(1, nil)
	bst.Insert(2, nil)
	bst.Insert(4, nil)
	bst.Insert(7, nil)
	bst.Insert(5, nil)
	fmt.Print("In Order: ")
	arr := bst.InOrderTraversal() // expect: 1, 2, 3, 4, 5, 6, 7
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i].Key, ", ")
	}
	fmt.Print("\nPre Order: ")
	arr = bst.PreOrderTraversal() // expect: 6, 3, 1, 2, 4, 5, 7
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i].Key, ", ")
	}
	fmt.Print("\nPost Order: ")
	arr = bst.PostOrderTraversal() // expect: 2, 1, 5, 4, 3, 7, 6
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i].Key, ", ")
	}
	fmt.Print("\nLevel Order: ")
	arr = bst.LevelOrderTraversal() // expect: 6, 3, 7, 1, 4, 2, 5
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i].Key, ", ")
	}
	fmt.Println("")
	fmt.Println("Contains 2?", bst.Contains(2)) // expect: true
	fmt.Println("Contains 4?", bst.Contains(4)) // expect: true
	fmt.Println("Contains 6?", bst.Contains(6)) // expect: true
	bst.Remove(2)
	bst.Remove(4)
	bst.Remove(6)
	fmt.Println("Contains 2?", bst.Contains(2)) // expect: false
	fmt.Println("Contains 4?", bst.Contains(4)) // expect: false
	fmt.Println("Contains 6?", bst.Contains(6)) // expect: false
	fmt.Print("In Order: ")
	arr = bst.InOrderTraversal() // expect: 1, 3, 5, 7
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i].Key, ", ")
	}
	fmt.Print("\nPre Order: ")
	arr = bst.PreOrderTraversal() // expect: 7, 3, 1, 5
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i].Key, ", ")
	}
	fmt.Print("\nPost Order: ")
	arr = bst.PostOrderTraversal() // expect: 1, 5, 3, 7
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i].Key, ", ")
	}
	fmt.Print("\nLevel Order: ")
	arr = bst.LevelOrderTraversal() // expect: 7, 3, 1, 5
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i].Key, ", ")
	}
	fmt.Println("")

	fmt.Println("Empty?", bst.Empty())   // expect: false
	fmt.Println("Root:", bst.Root().Key) // expect: Root: 7
	fmt.Println(bst.GetPair(3))          // expect: {3 <nil>}
	bst.UpdateVal(3, "three")
	fmt.Println(bst.GetPair(3))                                  // expect: {3 three}
	fmt.Println("Empty?", ds.NewBinarySearchTree[int]().Empty()) // expect: true

	fmt.Println("\n\n\n\t>>> Set <<<")
	set := ds.NewSet[int]()
	set.Insert(4)
	set.Insert(3)
	set.Insert(2)
	set.Insert(1)
	set.Insert(0)
	fmt.Println("Contains 1?", set.Contains(1)) // expect: true
	fmt.Println("Size:", set.Size())            // expect: 5
	fmt.Println(set.Values())                   // expect: 0, 1, 2, 3, 4 in any order
	set.Insert(1)
	fmt.Println("Contains 1?", set.Contains(1)) // expect: true
	fmt.Println("Size:", set.Size())            // expect: 5
	fmt.Println(set.Values())                   // expect: 0, 1, 2, 3, 4 in any order
	set.Remove(1)
	fmt.Println("Contains 1?", set.Contains(1)) // expect: false
	fmt.Println("Size:", set.Size())            // expect: 4
	fmt.Println(set.Values())                   // expect: 0, 2, 3, 4 in any order
	set.Remove(4)
	fmt.Println("Contains 4?", set.Contains(4)) // expect: false
	fmt.Println("Size:", set.Size())            // expect: 3
	fmt.Println(set.Values())                   // expect: 0, 2, 3 in any order

	fmt.Println("Empty?", set.Empty())              // expect: false
	fromSlice := ds.SetFromSlice([]int{1, 2, 3, 2}) // expect: The Key: 2 already exist on the Set
	fmt.Println("Size:", fromSlice.Size())          // expect: 3
	fmt.Println(fromSlice.Values())                 // expect: 1, 2, 3 in any order
	fmt.Println("Empty?", ds.NewSet[int]().Empty()) // expect: true

	fmt.Println("\n\n\n\t>>> Map <<<")
	hashmap := ds.NewMap[string, int]()
	hashmap.Set("e", 1)
	hashmap.Set("m", 2)
	hashmap.Set("a", 1)
	fmt.Println("Contains a?", hashmap.Contains("a")) // expect: true
	fmt.Println(hashmap.Get("a"))                     // expect: 1
	fmt.Println("Size:", hashmap.Size())              // expect: 3
	fmt.Println("Keys:", hashmap.Keys())              // expect: e, m, a in any order
	fmt.Println("Values:", hashmap.Values())          // expect: 1, 2, 1 in any order
	fmt.Println("Objects:", hashmap.Objects())        // expect: {e, 1}, {a, 1}, {m, 2} in any order
	hashmap.Set("d", 1)
	hashmap.Set("o", 1)
	hashmap.Set("r", 1)
	hashmap.Set("a", 2)
	fmt.Println(hashmap.Get("a"))              // expect: 2
	fmt.Println("Size:", hashmap.Size())       // expect: 6
	fmt.Println("Keys:", hashmap.Keys())       // expect: e, m, a, d, o, r in any order
	fmt.Println("Values:", hashmap.Values())   // expect: 1, 2, 1, 2, 1, 1 in any order
	fmt.Println("Objects:", hashmap.Objects()) // expect: {e, 1}, {a, 2}, {m, 2}, {d, 1}, {o, 1}, {r, 1} in any order
	hashmap.Remove("e")
	hashmap.Remove("m")
	fmt.Println("Contains m?", hashmap.Contains("m")) // expect: false
	fmt.Println("Size:", hashmap.Size())              // expect: 4
	fmt.Println("Keys:", hashmap.Keys())              // expect: a, d, o, r in any order
	fmt.Println("Values:", hashmap.Values())          // expect: 1, 2, 1, 1 in any order
	fmt.Println("Objects:", hashmap.Objects())        // expect: {a, 2}, {d, 1}, {o, 1}, {r, 1} in any order

	fmt.Println("Empty?", hashmap.Empty())                  // expect: false
	fmt.Println("Empty?", ds.NewMap[string, int]().Empty()) // expect: true

	fmt.Println("\n\n\n\t>>> MaxHeap <<<")
	maxHeap := ds.NewMaxHeap[int]()
	maxHeap.Insert(3)
	maxHeap.Insert(1)
	maxHeap.Insert(4)
	maxHeap.Insert(2)
	fmt.Println("Size:", maxHeap.Size())     // expect: 4
	fmt.Println("Sorted:", maxHeap.Sorted()) // expect: 4, 3, 2, 1
	for maxHeap.Empty() == false {
		fmt.Print(maxHeap.Peek(), ", ")
		maxHeap.RemoveTop()
	}
	// expect: 4, 3, 2, 1
	fmt.Println("Size:", maxHeap.Size()) // expect: 0

	fmt.Println("\n\n\n\t>>> MinHeap <<<")
	minHeap := ds.NewMinHeap[int]()
	minHeap.Insert(3)
	minHeap.Insert(1)
	minHeap.Insert(4)
	minHeap.Insert(2)
	fmt.Println("Size:", minHeap.Size())     // expect: 4
	fmt.Println("Sorted:", minHeap.Sorted()) // expect: 1, 2, 3, 4
	for minHeap.Empty() == false {
		fmt.Print(minHeap.Peek(), ", ")
		minHeap.RemoveTop()
	}
	// expect: 1, 2, 3, 4
	fmt.Println("Size:", minHeap.Size()) // expect: 0

	fmt.Println("\n\n\n\t>>> Priority Queue <<<")
	pq := ds.NewPriorityQueue[int, string]()
	pq.Enqueue(1, "a")
	pq.Enqueue(2, "b")
	pq.Enqueue(0, "c")
	pq.Enqueue(1, "d")
	fmt.Println("Size:", pq.Size()) // expect: 4
	for pq.Empty() == false {
		fmt.Print(pq.Front().Value, ", ")
		pq.Dequeue()
	}
	// expect: b, a, d, c
	fmt.Println("Size:", pq.Size()) // expect: 0

	fmt.Println("\n\n\n\t>>> Graph <<<")
	graph := ds.NewGraph[string]()
	graph.Insert(ds.Pair[string]{Key: "a", Value: nil}, []string{})
	graph.Insert(ds.Pair[string]{Key: "b", Value: nil}, []string{})
	graph.Insert(ds.Pair[string]{Key: "c", Value: nil}, []string{"a", "b"})
	graph.Insert(ds.Pair[string]{Key: "d", Value: nil}, []string{"b", "c"})
	graph.Insert(ds.Pair[string]{Key: "e", Value: nil}, []string{"a"})
	graph.Insert(ds.Pair[string]{Key: "f", Value: nil}, []string{})
	fmt.Println("Size:", graph.Size())         // expect: 6
	fmt.Println(graph.BreadthFirstTraversal()) // expect: a, c, e, b, d, f
	fmt.Println(graph.DepthFirstTraversal())   // expect: a, c, b, d, e, f
	graph.Remove("a")
	fmt.Println("Size:", graph.Size())                   // expect: 5
	fmt.Println(graph.BreadthFirstTraversal())           // expect: b, c, d, e, f
	fmt.Println(graph.DepthFirstTraversal())             // expect: b, c, d, e, f
	fmt.Println("Empty?", graph.Empty())                 // expect: false
	fmt.Println("Empty?", ds.NewGraph[string]().Empty()) // expect: true
}
