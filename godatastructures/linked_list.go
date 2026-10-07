package godatastructures

import "fmt"

type SingleLinkNode[T any] struct {
	Data T
	Next *SingleLinkNode[T]
}

type DoubleLinkNode[T any] struct {
	Data T
	Next *DoubleLinkNode[T]
	Prev *DoubleLinkNode[T]
}

type SinglyLinkedListInterface[T any] interface {
	InsertAfter(val T, prev *SingleLinkNode[T]) // create node with Data val after prevNode, increase size by 1. O(1)
	RemoveAfter(prev *SingleLinkNode[T])        // remove node after prevNode, decrease size by 1. O(1)
	InsertAtFront(val T)                        // create node with Data val at front of list, increase size by 1. O(1)
	RemoveAtFront()                             // remove node at front of list, decrease size by 1. O(1)
	InsertAtEnd(val T)                          // create node with Data val at end of list, increase size by 1. O(n)
	RemoveAtEnd()                               // remove node at end of list, decrease size by 1. O(n)
	Head() *SingleLinkNode[T]                   // return first node in list. O(1)
	Tail() *SingleLinkNode[T]                   // return last node in list. O(n)
	Empty() bool                                // returns whether list is empty. O(1)
	Size() int                                  // returns number of elements in list. O(1) or O(n) depending on implementation
	Print()                                     // prints out the SinglyLinkedList
}

type DoublyLinkedListInterface[T any] interface {
	InsertAfter(val T, prev *DoubleLinkNode[T])  // create node with Data val after prevNode, increase size by 1. O(1)
	InsertBefore(val T, next *DoubleLinkNode[T]) // create node with Data val before nextNode, increase size by 1. O(1)
	Remove(node *DoubleLinkNode[T])              // remove node, decrease size by 1. O(1)
	InsertAtFront(val T)                         // create node with Data val at front of list, increase size by 1. O(1)
	RemoveAtFront()                              // remove node at front of list, decrease size by 1. O(1)
	InsertAtEnd(val T)                           // create node with Data val at end of list, increase size by 1. O(1)
	RemoveAtEnd()                                // remove node at end of list, decrease size by 1. O(1)
	Head() *DoubleLinkNode[T]                    // return first node in list. O(1)
	Tail() *DoubleLinkNode[T]                    // return last node in list. O(1)
	Empty() bool                                 // returns whether list is empty. O(1)
	Size() int                                   // returns number of elements in list. O(1) or O(n) depending on implementation
	Print()                                      // prints out the DoublyLinkedList
}

// SinglyLinkedList should implement SinglyLinkedListInterface; write a constructor & methods to complete it
type SinglyLinkedList[T any] struct {
	head *SingleLinkNode[T]
	size int // you may choose whether to uncomment this line
}

// ModifiedSinglyLinkedList should implement SinglyLinkedListInterface; write a constructor & methods to complete it
type ModifiedSinglyLinkedList[T any] struct {
	head *SingleLinkNode[T]
	tail *SingleLinkNode[T]
	size int // you may choose whether to uncomment this line
}

// DoublyLinkedList should implement DoublyLinkedListInterface; write a constructor & methods to complete it
type DoublyLinkedList[T any] struct {
	head *DoubleLinkNode[T]
	tail *DoubleLinkNode[T]
	size int // you may choose whether to uncomment this line
}

/*      Methods: SinglyLinkedList struct      */
func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{nil, 0}
}

// create node with Data val after prevNode, increase size by 1. O(1)
func (sll *SinglyLinkedList[T]) InsertAfter(val T, prevNode *SingleLinkNode[T]) {
	newNode := &SingleLinkNode[T]{Data: val, Next: prevNode.Next}
	prevNode.Next = newNode
	sll.size++
}

// remove node after prevNode, decrease size by 1. O(1)
func (sll *SinglyLinkedList[T]) RemoveAfter(prevNode *SingleLinkNode[T]) {
	// if the prevNode is the tail:
	if prevNode.Next == nil {
		return
	}

	fmt.Println("SinglyLinkedList RemoveAfter: removed", prevNode.Next.Data)

	prevNode.Next = prevNode.Next.Next
	sll.size--
}

// create node with Data val at front of list, increase size by 1. O(1)
func (sll *SinglyLinkedList[T]) InsertAtFront(val T) {
	newNode := &SingleLinkNode[T]{Data: val, Next: sll.head}
	sll.head = newNode
	sll.size++
}

// remove node at front of list, decrease size by 1. O(1)
func (sll *SinglyLinkedList[T]) RemoveAtFront() {
	if sll.head == nil {
		fmt.Println("SinglyLinkedList is empty: nil head")
		return
	}

	fmt.Println("SinglyLinkedList RemoveAtFront: removed", sll.head.Data)

	sll.head = sll.head.Next
	sll.size--
}

// create node with Data val at end of list, increase size by 1. O(n)
func (sll *SinglyLinkedList[T]) InsertAtEnd(val T) {
	newNode := &SingleLinkNode[T]{Data: val, Next: nil}

	if sll.Head() == nil {
		sll.head = newNode
	} else {
		lastNode := sll.Tail()
		lastNode.Next = newNode
	}
	sll.size++
}

// remove node at end of list, decrease size by 1. O(n)
func (sll *SinglyLinkedList[T]) RemoveAtEnd() {
	if sll.head == nil {
		fmt.Println("SinglyLinkedList is empty: nil head")
		return
	}

	// if the currNode (head) is the tail (linked list has only one node):
	if sll.head.Next == nil {
		sll.RemoveAtFront()
		return
	}

	currNode := sll.head
	for currNode.Next.Next != nil {
		currNode = currNode.Next
	}
	fmt.Println("SinglyLinkedList RemoveAtEnd: removed", currNode.Next.Data)
	currNode.Next = nil
	sll.size--
}

// return first node in list. O(1)
func (sll *SinglyLinkedList[T]) Head() *SingleLinkNode[T] {
	return sll.head
}

// return last node in list. O(n)
func (sll *SinglyLinkedList[T]) Tail() *SingleLinkNode[T] {
	if sll.head == nil {
		return nil
	}

	currNode := sll.head
	for currNode.Next != nil {
		currNode = currNode.Next
	}
	return currNode
}

// returns whether list is empty. O(1)
func (sll *SinglyLinkedList[T]) Empty() bool {
	return sll.size == 0
}

// returns number of elements in list. O(1) or O(n) depending on implementation
func (sll *SinglyLinkedList[T]) Size() int {
	return sll.size
}

// prints out the SinglyLinkedList
func (sll *SinglyLinkedList[T]) Print() {
	curr := sll.head
	for curr != nil {
		fmt.Print(curr.Data, " -> ")
		curr = curr.Next
	}
	fmt.Println(nil)
}

/*      Methods: DoublyLinkedListInterface struct      */
func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{nil, nil, 0}
}

// create node with Data val after prevNode, increase size by 1. O(1)
func (dll *DoublyLinkedList[T]) InsertAfter(val T, prevNode *DoubleLinkNode[T]) {
	newNode := &DoubleLinkNode[T]{Data: val, Next: prevNode.Next, Prev: prevNode}
	if prevNode.Next == nil {
		dll.tail = newNode
	} else {
		prevNode.Next.Prev = newNode
	}
	prevNode.Next = newNode
	dll.size++
}

// create node with Data val before nextNode, increase size by 1. O(1)
func (dll *DoublyLinkedList[T]) InsertBefore(val T, nextNode *DoubleLinkNode[T]) {
	newNode := &DoubleLinkNode[T]{Data: val, Next: nextNode, Prev: nextNode.Prev}
	if nextNode.Prev == nil {
		dll.head = newNode
	} else {
		nextNode.Prev.Next = newNode
	}
	nextNode.Prev = newNode
	dll.size++
}

// remove node, decrease size by 1. O(1)
func (dll *DoublyLinkedList[T]) Remove(node *DoubleLinkNode[T]) {
	fmt.Println("DoublyLinkedList Remove: removed", node.Data)

	if dll.head == node {
		dll.head = node.Next
		//cutting both connections
		if node.Next != nil {
			node.Next.Prev = nil
		} else { // node was the only node in the list
			dll.tail = nil
		}
		node.Next = nil
	} else if dll.tail == node {
		dll.tail = node.Prev
		//cutting both connections
		node.Prev.Next = nil
		node.Prev = nil
	} else {
		//linking Prev node with Next node, also their Prevs and Nexts
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		//cutting current node out of the chain
		node.Prev = nil
		node.Next = nil
	}
	dll.size--
}

// create node with Data val at front of list, increase size by 1. O(1)
func (dll *DoublyLinkedList[T]) InsertAtFront(val T) {
	newNode := &DoubleLinkNode[T]{Data: val, Next: dll.head, Prev: nil}
	if dll.head == nil {
		dll.tail = newNode
	} else {
		dll.head.Prev = newNode
	}
	dll.head = newNode
	dll.size++
}

// remove node at front of list, decrease size by 1. O(1)
func (dll *DoublyLinkedList[T]) RemoveAtFront() {
	// dll.Remove(dll.head)
	if dll.head == nil {
		fmt.Println("DoublyLinkedList is empty: nil head")
		return
	}

	fmt.Println("DoublyLinkedList RemoveAtFront: removed", dll.head.Data)

	if dll.head.Next != nil {
		dll.head.Next.Prev = nil
	} else { // the only node is being removed
		dll.tail = nil
	}

	oldHead := dll.head
	dll.head = oldHead.Next
	oldHead.Next = nil
	oldHead.Prev = nil
	dll.size--
}

// create node with Data val at end of list, increase size by 1. O(1)
func (dll *DoublyLinkedList[T]) InsertAtEnd(val T) {
	newNode := &DoubleLinkNode[T]{Data: val, Next: nil, Prev: dll.tail}
	if dll.tail == nil {
		dll.head = newNode
	} else {
		dll.tail.Next = newNode
	}
	dll.tail = newNode
	dll.size++
}

// remove node at end of list, decrease size by 1. O(1)
func (dll *DoublyLinkedList[T]) RemoveAtEnd() {
	// dll.Remove(dll.tail)
	if dll.tail == nil {
		fmt.Println("DoublyLinkedList is empty: nil tail")
		return
	}

	fmt.Println("DoublyLinkedList RemoveAtEnd: removed", dll.tail.Data)

	if dll.tail.Prev != nil {
		dll.tail.Prev.Next = nil
	} else { // the only node is being removed
		dll.head = nil
	}

	oldTail := dll.tail
	dll.tail = oldTail.Prev
	oldTail.Prev = nil
	oldTail.Next = nil
	dll.size--
}

// return first node in list. O(1)
func (dll *DoublyLinkedList[T]) Head() *DoubleLinkNode[T] {
	return dll.head
}

// return last node in list. O(1)
func (dll *DoublyLinkedList[T]) Tail() *DoubleLinkNode[T] {
	return dll.tail
}

// returns whether list is empty. O(1)
func (dll *DoublyLinkedList[T]) Empty() bool {
	return dll.size == 0
}

// returns number of elements in list. O(1) or O(n) depending on implementation
func (dll *DoublyLinkedList[T]) Size() int {
	return dll.size
}

// prints out the DoublyLinkedList
func (dll *DoublyLinkedList[T]) Print() {
	curr := dll.head
	for curr != nil {
		fmt.Print(curr.Data, " -> ")
		curr = curr.Next
	}
	fmt.Println(nil)
}

/*      Methods: ModifiedSinglyLinkedList struct      */
func NewModifiedSinglyLinkedList[T any]() *ModifiedSinglyLinkedList[T] {
	return &ModifiedSinglyLinkedList[T]{nil, nil, 0}
}

// create node with Data val after prevNode, increase size by 1. O(1)
func (msll *ModifiedSinglyLinkedList[T]) InsertAfter(val T, prevNode *SingleLinkNode[T]) {
	newNode := &SingleLinkNode[T]{Data: val, Next: prevNode.Next}
	prevNode.Next = newNode
	// if the prevNode was the tail:
	if msll.tail == prevNode {
		msll.tail = newNode
	}
	msll.size++
}

// remove node after prevNode, decrease size by 1. O(1)
func (msll *ModifiedSinglyLinkedList[T]) RemoveAfter(prevNode *SingleLinkNode[T]) {
	// if the prevNode is the tail:
	if prevNode.Next == nil {
		return
	}

	fmt.Println("ModifiedSinglyLinkedList RemoveAfter: removed", prevNode.Next.Data)

	// if the removed node was the tail:
	if prevNode.Next == msll.tail {
		msll.tail = prevNode
	}
	prevNode.Next = prevNode.Next.Next
	msll.size--
}

// create node with Data val at front of list, increase size by 1. O(1)
func (msll *ModifiedSinglyLinkedList[T]) InsertAtFront(val T) {
	newNode := &SingleLinkNode[T]{Data: val, Next: msll.head}
	if msll.head == nil {
		msll.tail = newNode
	}
	msll.head = newNode
	msll.size++
}

// remove node at front of list, decrease size by 1. O(1)
func (msll *ModifiedSinglyLinkedList[T]) RemoveAtFront() {
	if msll.head == nil {
		fmt.Println("ModifiedSinglyLinkedList is empty: nil head")
		return
	}

	fmt.Println("ModifiedSinglyLinkedList RemoveAtFront: removed", msll.head.Data)

	// if the only node is being removed:
	if msll.head.Next == nil {
		msll.tail = nil
	}
	msll.head = msll.head.Next
	msll.size--
}

// create node with Data val at end of list, increase size by 1. O(1)
func (msll *ModifiedSinglyLinkedList[T]) InsertAtEnd(val T) {
	newNode := &SingleLinkNode[T]{Data: val, Next: nil}

	if msll.tail == nil {
		msll.head = newNode
	} else {
		msll.tail.Next = newNode
	}
	msll.tail = newNode
	msll.size++
}

// remove node at end of list, decrease size by 1. O(n)
func (msll *ModifiedSinglyLinkedList[T]) RemoveAtEnd() {
	if msll.head == nil {
		fmt.Println("ModifiedSinglyLinkedList is empty: nil head")
		return
	}

	// if the currNode (head) is the tail (linked list has only one node):
	if msll.head.Next == nil {
		msll.RemoveAtFront()
		return
	}

	currNode := msll.head
	for currNode.Next.Next != nil {
		currNode = currNode.Next
	}
	fmt.Println("ModifiedSinglyLinkedList RemoveAtEnd: removed", currNode.Next.Data)
	currNode.Next = nil
	msll.tail = currNode
	msll.size--
}

// return first node in list. O(1)
func (msll *ModifiedSinglyLinkedList[T]) Head() *SingleLinkNode[T] {
	return msll.head
}

// return last node in list. O(1)
func (msll *ModifiedSinglyLinkedList[T]) Tail() *SingleLinkNode[T] {
	return msll.tail
}

// returns whether list is empty. O(1)
func (msll *ModifiedSinglyLinkedList[T]) Empty() bool {
	return msll.size == 0
}

// returns number of elements in list. O(1) or O(n) depending on implementation
func (msll *ModifiedSinglyLinkedList[T]) Size() int {
	return msll.size
}

// prints out the ModifiedSinglyLinkedList
func (msll *ModifiedSinglyLinkedList[T]) Print() {
	curr := msll.head
	for curr != nil {
		fmt.Print(curr.Data, " -> ")
		curr = curr.Next
	}
	fmt.Println(nil)
}
