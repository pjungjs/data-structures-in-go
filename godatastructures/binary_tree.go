package godatastructures

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type BinaryTreeNode[T constraints.Ordered] struct {
	Key   T
	Value any
	Left  *BinaryTreeNode[T]
	Right *BinaryTreeNode[T]
}

type BinaryTreeInterface[T constraints.Ordered] interface {
	Root() *BinaryTreeNode[T]       // returns node at root of tree. O(1)
	GetPair(key T) Pair[T]          // returns node's pair (Key & Value) by key
	UpdateVal(key T, val any)       // update node's val by key
	Insert(key T, val any)          // inserts node with key, val (can be nil), increases size by 1. O(logn) on average, O(n) worst case
	Contains(key T) bool            // checks if tree contains key. O(logn) on average, O(n) worst case
	Remove(key T)                   // removes node with key, decreases size by 1. O(logn) on average, O(n) worst case
	Empty() bool                    // returns whether tree is empty. O(1)
	InOrderTraversal() []Pair[T]    // returns data in tree ordered by processing left, current, right. O(n)
	PreOrderTraversal() []Pair[T]   // returns data in tree ordered by processing current, left, right. O(n)
	PostOrderTraversal() []Pair[T]  // returns data in tree ordered by processing left, right, current. O(n)
	LevelOrderTraversal() []Pair[T] // returns data in tree ordered by processing each level left to right. O(n)
}

// BinarySearchTree should implement BinaryTreeInterface; write a constructor & methods to complete it
type BinarySearchTree[T constraints.Ordered] struct {
	root *BinaryTreeNode[T]
}

/*      Methods: BinarySearchTree struct      */
func NewBinarySearchTree[T constraints.Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{nil}
}

// returns node at root of tree. O(1)
func (bst *BinarySearchTree[T]) Root() *BinaryTreeNode[T] {
	return bst.root
}

// returns node's pair (Key & Value) by key
func (bst *BinarySearchTree[T]) GetPair(key T) Pair[T] {
	pair := Pair[T]{}

	currNode := bst.root
	for currNode != nil {
		if key < currNode.Key {
			currNode = currNode.Left
		} else if key > currNode.Key {
			currNode = currNode.Right
		} else {
			pair = Pair[T]{currNode.Key, currNode.Value}
			break
		}
	}
	return pair
}

// update node's val by key
func (bst *BinarySearchTree[T]) UpdateVal(key T, val any) {
	currNode := bst.root
	for currNode != nil {
		if key < currNode.Key {
			currNode = currNode.Left
		} else if key > currNode.Key {
			currNode = currNode.Right
		} else {
			currNode.Value = val
			return
		}
	}
}

// inserts node with key, val (can be nil), increases size by 1. O(logn) on average, O(n) worst case
func (bst *BinarySearchTree[T]) Insert(key T, val any) {
	newNode := &BinaryTreeNode[T]{Key: key, Value: val, Left: nil, Right: nil}

	if bst.root == nil {
		bst.root = newNode
		return
	}

	currNode := bst.root
	for currNode != nil {
		if key < currNode.Key {
			if currNode.Left == nil {
				currNode.Left = newNode
				return
			}
			currNode = currNode.Left
		} else if key > currNode.Key {
			if currNode.Right == nil {
				currNode.Right = newNode
				return
			}
			currNode = currNode.Right
		} else {
			//edge case to handle duplicated key
			fmt.Printf("Key %v already exists in the Binary Search Tree\n", key)
			return
		}
	}
}

// checks if tree contains key. O(logn) on average, O(n) worst case
func (bst *BinarySearchTree[T]) Contains(key T) bool {
	currNode := bst.root
	for currNode != nil {
		if key < currNode.Key {
			currNode = currNode.Left
		} else if key > currNode.Key {
			currNode = currNode.Right
		} else {
			return true
		}
	}
	return false
}

// removes node with key, throws error if does not exist, decreases size by 1. O(logn) on average, O(n) worst case
func (bst *BinarySearchTree[T]) Remove(key T) {
	/* 3 possible cases:
	   1. key node is a leaf node: remove it directly
	   2. key node has only one child: replace it with its child
	   3. key node has both two children:
	     find and replace it with (either or):
	     a. successor (smallest key in right subtree)
	     OR
	     b. predecessor (largest key in left subtree)
	*/

	// if the root is nil OR if the tree doesn't contain the key, print an error message
	if bst.root == nil || !bst.Contains(key) {
		fmt.Printf("Key %v does not exist in the Binary Search Tree\n", key)
		return
	}

	var parentNode *BinaryTreeNode[T]
	currNode := bst.root

	//find the node with the key and its parent
	for currNode.Key != key {
		parentNode = currNode
		if key < currNode.Key {
			currNode = currNode.Left
		} else {
			currNode = currNode.Right
		}
	}

	// case 1. key node is a leaf node: remove it directly
	if currNode.Left == nil && currNode.Right == nil {
		if currNode == bst.root { // node is root
			bst.root = nil
		} else if parentNode.Left == currNode { // node is left child
			parentNode.Left = nil
		} else { // node is right child
			parentNode.Right = nil
		}
		return
	}

	// case 2. key node has only one child: replace it with its child
	if currNode.Left == nil { // node has right child
		if currNode == bst.root { // node is root
			bst.root = currNode.Right
		} else if parentNode.Left == currNode { // node is left child
			parentNode.Left = currNode.Right
		} else { // node is right child
			parentNode.Right = currNode.Right
		}
		return
	} else if currNode.Right == nil { // node has left child
		if currNode == bst.root { // node is root
			bst.root = currNode.Left
		} else if parentNode.Left == currNode { // node is left child
			parentNode.Left = currNode.Left
		} else { // node is right child
			parentNode.Right = currNode.Left
		}
		return
	}

	// case 3. key node has both two children
	//find and replace it with successor (smallest key in right subtree) node
	successorParent := currNode
	successor := currNode.Right

	//find the smallest key in right subtree
	for successor.Left != nil {
		successorParent = successor
		successor = successor.Left
	}

	//if the successor is not the direct right child of the currNode
	if successorParent != currNode {
		successorParent.Left = successor.Right
		successor.Right = currNode.Right
	}

	//replace the node with the successor
	//connect the successor with its parent
	if parentNode == nil {
		bst.root = successor
	} else if parentNode.Left == currNode {
		parentNode.Left = successor
	} else {
		parentNode.Right = successor
	}

	// connect the successor with its left child
	successor.Left = currNode.Left

	/*
	     - parentNode: parent of currNode
	     - currNode: node that needs to be removed
	     - successorParent: parent of successor
	     - successor: node that will be replacing currNode's place

	   Example:

	   	    10
	   	   /  \
	   	  5    15
	   	 / \   / \
	   	3   8 12  20
	   	     \
	   	      9
	*/
}

// returns whether tree is empty. O(1)
func (bst *BinarySearchTree[T]) Empty() bool {
	return bst.root == nil
}

// Depth-First Search (DFS) - Stack
// returns data in tree ordered by processing left, current, right. O(n)
func (bst *BinarySearchTree[T]) InOrderTraversal() []Pair[T] {
	var inOrderedVals []Pair[T]

	//using slices:
	currNode := bst.root
	stack := make([]*BinaryTreeNode[T], 0)

	for currNode != nil || len(stack) > 0 {
		for currNode != nil {
			stack = append(stack, currNode)
			currNode = currNode.Left
		}

		currNode = stack[len(stack)-1] //top of the stack
		stack = stack[:len(stack)-1]   //removing last element of the stack
		inOrderedVals = append(inOrderedVals, Pair[T]{currNode.Key, currNode.Value})
		currNode = currNode.Right
	}
	return inOrderedVals
}

// returns data in tree ordered by processing current, left, right. O(n)
func (bst *BinarySearchTree[T]) PreOrderTraversal() []Pair[T] {
	var preOrderedVals []Pair[T]

	currNode := bst.root
	stack := make([]*BinaryTreeNode[T], 0)

	for currNode != nil || len(stack) > 0 {
		if currNode != nil {
			preOrderedVals = append(preOrderedVals, Pair[T]{currNode.Key, currNode.Value})
			if currNode.Right != nil {
				stack = append(stack, currNode.Right)
			}
			currNode = currNode.Left
		} else {
			currNode = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
	return preOrderedVals
}

// returns data in tree ordered by processing left, right, current. O(n)
func (bst *BinarySearchTree[T]) PostOrderTraversal() []Pair[T] {
	/*different ways of doing post order traversal:
	  1. do current, right, left, then reverse it at the end (reversed modified-preorder)
	  2. have two stacks tracking the stack and the ones that you visited before
	  3. track stack with slice, but prevNode, currNode and lastStack using var
	*/
	var (
		postOrderedVals []Pair[T]
		currNode        *BinaryTreeNode[T] = bst.root
		prevNode        *BinaryTreeNode[T]
		stack           []*BinaryTreeNode[T]
	)

	for currNode != nil || len(stack) > 0 {
		for currNode != nil {
			//go to the leftmost leaf node and add it to the stack along the way
			stack = append(stack, currNode)
			currNode = currNode.Left
		}

		lastStack := stack[len(stack)-1]
		if lastStack.Right == nil || lastStack.Right == prevNode {
			//if the Right subtree doesn't exist or has been visited
			postOrderedVals = append(postOrderedVals, Pair[T]{lastStack.Key, lastStack.Value})
			stack = stack[:len(stack)-1]
			prevNode = lastStack
		} else {
			currNode = lastStack.Right
		}
	}
	return postOrderedVals
}

// Breadth-First Search (BFS) - Queue
// returns data in tree ordered by processing each level left to right. O(n)
func (bst *BinarySearchTree[T]) LevelOrderTraversal() []Pair[T] {
	var levelOrderedVals []Pair[T]

	//using slices:
	queue := []*BinaryTreeNode[T]{bst.root}

	if bst.root != nil {
		for len(queue) > 0 {
			currNode := queue[0] //first element of the queue
			queue = queue[1:]    //removes the first element of the queue
			levelOrderedVals = append(levelOrderedVals, Pair[T]{currNode.Key, currNode.Value})

			if currNode.Left != nil {
				queue = append(queue, currNode.Left)
			}
			if currNode.Right != nil {
				queue = append(queue, currNode.Right)
			}
		}
	}
	return levelOrderedVals
}
