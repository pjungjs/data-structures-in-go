# Data Structures in Go

Common data structures implemented from scratch in Go (generics), originally written by hand as a learning project back in early 2024.

Each structure is built on top of the earlier ones where it makes sense. For example, `Stack` and `Queue` use a linked list, `Set` and `Map` use a hash table whose buckets are binary search trees, and `PriorityQueue` uses a max heap.

## What's included

All code lives in [`godatastructures/`](./godatastructures).

| Structure | File | Notes |
|---|---|---|
| `SinglyLinkedList` | `linked_list.go` | Head only. O(1) at the front, O(n) at the end |
| `ModifiedSinglyLinkedList` | `linked_list.go` | Adds a tail pointer. O(1) `InsertAtEnd` and `Tail`, O(n) `RemoveAtEnd` |
| `DoublyLinkedList` | `linked_list.go` | Head and tail pointers. O(1) insert and remove at both ends and by node |
| `Stack`, `AlternateStack` | `stack.go` | Linked list based and slice based. O(1) `Push`, `Pop`, `Top` (amortized for the slice `Push`) |
| `Queue`, `AlternateQueue` | `queue.go` | Linked list based and slice based. O(1) `Dequeue`, `Front`. `Enqueue` is O(n) for `Queue`, amortized O(1) for `AlternateQueue` |
| `Deque` | `deque.go` | Built on the doubly linked list. O(1) at both ends |
| `BinarySearchTree` | `binary_tree.go` | Unbalanced, with in/pre/post/level-order traversals. O(log n) average and O(n) worst case per operation, O(n) traversals |
| `Set` | `set.go` | Hash table (FNV) with BST buckets, resizes at a max fill ratio. O(1) average and O(n) worst case, O(n) `Values` |
| `Map` | `map.go` | Same design as `Set`, with key-value pairs. O(1) average and O(n) worst case, O(n) `Keys`, `Values`, `Objects` |
| `MaxHeap`, `MinHeap` | `heap.go` | Array based. O(log n) `Insert`, `RemoveTop`, O(1) `Peek`, O(n log n) `Sorted` |
| `PriorityQueue` | `priority_queue.go` | Built on a max heap. O(log n) `Enqueue`, `Dequeue`, O(1) `Front` |
| `Graph` | `graph.go` | Undirected, unweighted adjacency list. O(k) `Insert` with k neighbors, O(V) `Remove`, O(V + E log V) DFS and BFS |

Space is O(n) for every structure (O(V + E) for `Graph`), where `V` and `E` are the number of nodes and edges. `Size` and `Empty` are O(1) everywhere.

Every structure has a matching `...Interface` type documenting its methods and their expected time complexity.

## Usage

```go
package main

import (
	"fmt"

	ds "github.com/pjungjs/data-structures-in-go/godatastructures"
)

func main() {
	stack := ds.NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Top()) // 2

	m := ds.NewMap[string, int]()
	m.Set("a", 1)
	fmt.Println(m.Get("a")) // 1
}
```

## Running the demo

`main.go` walks through every structure and prints the result, with the expected output in the comments:

```sh
go run .
```

## Running the tests

```sh
go test ./...
```

The linked lists have unit tests in `godatastructures/linked_list_test.go`. The other structures are exercised by the demo in `main.go`.

## Behavior notes

- Mutating methods print a short message (for example `removed 3`) as they run, so the demo output is easy to follow.
- Operations on an empty structure (such as `Pop` or `Dequeue`) print a message and return, and `Top`, `Front`, `Peek` and similar methods return the zero value instead of panicking.
- `Set`, `Map` and `Graph` return their elements in hash order, except that graph traversals visit neighbors in ascending order so results are deterministic.

## License

[MIT](./LICENSE)
