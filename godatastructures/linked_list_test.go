// cd godatastructures/
// go test
package godatastructures

import (
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

/*      Helper Functions      */

// runs f and returns everything it printed to stdout
func captureOutput(f func()) string {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = oldStdout
	out, _ := io.ReadAll(r)
	return string(out)
}

// returns the values of a singly linked list, from head to tail
func singleValues(head *SingleLinkNode[int]) []int {
	vals := []int{}
	for node := head; node != nil; node = node.Next {
		vals = append(vals, node.Data)
	}
	return vals
}

// returns the values of a doubly linked list from head to tail,
// and checks that every Prev pointer is consistent with Next
func doubleValues(t *testing.T, dll *DoublyLinkedList[int]) []int {
	t.Helper()
	vals := []int{}
	var prev *DoubleLinkNode[int]
	for node := dll.Head(); node != nil; node = node.Next {
		if node.Prev != prev {
			t.Errorf("Prev of node %d is wrong", node.Data)
		}
		vals = append(vals, node.Data)
		prev = node
	}
	if dll.Tail() != prev {
		t.Errorf("Tail is not the last node of the list")
	}
	return vals
}

func checkInts(t *testing.T, name string, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s: expected %v \t got %v", name, want, got)
	}
}

func checkInt(t *testing.T, name string, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("%s: expected %d \t got %d", name, want, got)
	}
}

func checkBool(t *testing.T, name string, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("%s: expected %t \t got %t", name, want, got)
	}
}

func checkOutput(t *testing.T, name, got, want string) {
	t.Helper()
	if !strings.Contains(got, want) {
		t.Errorf("%s: expected output to contain %q \t got %q", name, want, got)
	}
}

/*      Tests: SinglyLinkedList      */

func TestSinglyLinkedList(t *testing.T) {
	sll := NewSinglyLinkedList[int]()

	//empty list
	checkBool(t, "Empty (new)", sll.Empty(), true)
	checkInt(t, "Size (new)", sll.Size(), 0)
	if sll.Head() != nil || sll.Tail() != nil {
		t.Errorf("Head and Tail of an empty list should be nil")
	}

	sll.InsertAtEnd(11)                       //insert at the end of an empty list
	sll.InsertAfter(12, sll.Head())           //insert after 11
	sll.InsertAtEnd(13)                       //insert at the end
	sll.InsertAfter(14, sll.Head().Next.Next) //insert after 13
	sll.InsertAtEnd(15)                       //insert at the end
	sll.InsertAtFront(10)                     //insert at the beginning

	checkInt(t, "Head", sll.Head().Data, 10)
	checkInt(t, "Tail", sll.Tail().Data, 15)
	checkBool(t, "Empty", sll.Empty(), false)
	checkInt(t, "Size", sll.Size(), 6)
	checkInts(t, "values", singleValues(sll.Head()), []int{10, 11, 12, 13, 14, 15})

	out := captureOutput(sll.Print)
	checkOutput(t, "Print", out, "10 -> 11 -> 12 -> 13 -> 14 -> 15 -> <nil>")

	out = captureOutput(sll.RemoveAtFront) //removing 10
	checkOutput(t, "RemoveAtFront", out, "removed 10")
	out = captureOutput(func() { sll.RemoveAfter(sll.Head().Next.Next) }) //removing 14
	checkOutput(t, "RemoveAfter", out, "removed 14")
	out = captureOutput(sll.RemoveAtEnd) //removing 15
	checkOutput(t, "RemoveAtEnd", out, "removed 15")

	checkInt(t, "Head", sll.Head().Data, 11)
	checkInt(t, "Tail", sll.Tail().Data, 13)
	checkInt(t, "Size", sll.Size(), 3)
	checkInts(t, "values", singleValues(sll.Head()), []int{11, 12, 13})

	//RemoveAfter on the tail does nothing
	captureOutput(func() { sll.RemoveAfter(sll.Tail()) })
	checkInt(t, "Size (RemoveAfter tail)", sll.Size(), 3)
	checkInts(t, "values (RemoveAfter tail)", singleValues(sll.Head()), []int{11, 12, 13})

	//InsertAfter the tail
	sll.InsertAfter(14, sll.Tail())
	checkInt(t, "Tail (InsertAfter tail)", sll.Tail().Data, 14)
	checkInt(t, "Size (InsertAfter tail)", sll.Size(), 4)

	//removing everything, one node at a time
	captureOutput(func() {
		sll.RemoveAtEnd() //removing 14
		sll.RemoveAtEnd() //removing 13
		sll.RemoveAtEnd() //removing 12
		sll.RemoveAtEnd() //removing 11 (the only node)
	})
	checkBool(t, "Empty (all removed)", sll.Empty(), true)
	checkInt(t, "Size (all removed)", sll.Size(), 0)
	if sll.Head() != nil || sll.Tail() != nil {
		t.Errorf("Head and Tail should be nil after removing every node")
	}

	//removing from an empty list does not panic
	out = captureOutput(sll.RemoveAtFront)
	checkOutput(t, "RemoveAtFront (empty)", out, "is empty")
	out = captureOutput(sll.RemoveAtEnd)
	checkOutput(t, "RemoveAtEnd (empty)", out, "is empty")
	checkInt(t, "Size (empty removes)", sll.Size(), 0)

	//the list is still usable after being emptied
	sll.InsertAtFront(1)
	sll.InsertAtEnd(2)
	checkInts(t, "values (reuse)", singleValues(sll.Head()), []int{1, 2})

	out = captureOutput(NewSinglyLinkedList[int]().Print)
	checkOutput(t, "Print (empty)", out, "<nil>")
}

/*      Tests: ModifiedSinglyLinkedList      */

func TestModifiedSinglyLinkedList(t *testing.T) {
	msll := NewModifiedSinglyLinkedList[int]()

	//empty list
	checkBool(t, "Empty (new)", msll.Empty(), true)
	checkInt(t, "Size (new)", msll.Size(), 0)
	if msll.Head() != nil || msll.Tail() != nil {
		t.Errorf("Head and Tail of an empty list should be nil")
	}

	msll.InsertAtEnd(11) //insert at the end of an empty list
	checkInt(t, "Head (first node)", msll.Head().Data, 11)
	checkInt(t, "Tail (first node)", msll.Tail().Data, 11)
	msll.InsertAfter(12, msll.Head()) //insert after 11 (the tail)
	checkInt(t, "Tail (InsertAfter tail)", msll.Tail().Data, 12)
	msll.InsertAtEnd(14)                   //insert at the end
	msll.InsertAfter(13, msll.Head().Next) //insert after 12 (not the tail)
	msll.InsertAtEnd(15)                   //insert at the end
	msll.InsertAtFront(10)                 //insert at the beginning

	checkInt(t, "Head", msll.Head().Data, 10)
	checkInt(t, "Tail", msll.Tail().Data, 15)
	checkBool(t, "Empty", msll.Empty(), false)
	checkInt(t, "Size", msll.Size(), 6)
	checkInts(t, "values", singleValues(msll.Head()), []int{10, 11, 12, 13, 14, 15})

	out := captureOutput(msll.Print)
	checkOutput(t, "Print", out, "10 -> 11 -> 12 -> 13 -> 14 -> 15 -> <nil>")

	out = captureOutput(msll.RemoveAtFront) //removing 10
	checkOutput(t, "RemoveAtFront", out, "removed 10")
	out = captureOutput(func() { msll.RemoveAfter(msll.Head().Next.Next) }) //removing 14
	checkOutput(t, "RemoveAfter", out, "removed 14")
	out = captureOutput(msll.RemoveAtEnd) //removing 15
	checkOutput(t, "RemoveAtEnd", out, "removed 15")

	checkInt(t, "Head", msll.Head().Data, 11)
	checkInt(t, "Tail", msll.Tail().Data, 13)
	checkInt(t, "Size", msll.Size(), 3)
	checkInts(t, "values", singleValues(msll.Head()), []int{11, 12, 13})

	//RemoveAfter on the tail does nothing
	captureOutput(func() { msll.RemoveAfter(msll.Tail()) })
	checkInt(t, "Size (RemoveAfter tail)", msll.Size(), 3)

	//RemoveAfter the node before the tail moves the tail back
	captureOutput(func() { msll.RemoveAfter(msll.Head().Next) }) //removing 13
	checkInt(t, "Tail (RemoveAfter before tail)", msll.Tail().Data, 12)
	checkInts(t, "values (RemoveAfter before tail)", singleValues(msll.Head()), []int{11, 12})

	//removing everything
	captureOutput(func() {
		msll.RemoveAtEnd() //removing 12
		msll.RemoveAtEnd() //removing 11 (the only node)
	})
	checkBool(t, "Empty (all removed)", msll.Empty(), true)
	checkInt(t, "Size (all removed)", msll.Size(), 0)
	if msll.Head() != nil || msll.Tail() != nil {
		t.Errorf("Head and Tail should be nil after removing every node")
	}

	//removing from an empty list does not panic
	out = captureOutput(msll.RemoveAtFront)
	checkOutput(t, "RemoveAtFront (empty)", out, "is empty")
	out = captureOutput(msll.RemoveAtEnd)
	checkOutput(t, "RemoveAtEnd (empty)", out, "is empty")

	//the list is still usable after being emptied
	msll.InsertAtFront(1)
	checkInt(t, "Tail (InsertAtFront empty)", msll.Tail().Data, 1)
	msll.InsertAtEnd(2)
	checkInts(t, "values (reuse)", singleValues(msll.Head()), []int{1, 2})
	checkInt(t, "Tail (reuse)", msll.Tail().Data, 2)

	//RemoveAtFront of the only node clears the tail
	captureOutput(func() {
		msll.RemoveAtFront()
		msll.RemoveAtFront()
	})
	if msll.Tail() != nil {
		t.Errorf("Tail should be nil after removing the only node")
	}
}

/*      Tests: DoublyLinkedList      */

func TestDoublyLinkedList(t *testing.T) {
	dll := NewDoublyLinkedList[int]()

	//empty list
	checkBool(t, "Empty (new)", dll.Empty(), true)
	checkInt(t, "Size (new)", dll.Size(), 0)
	if dll.Head() != nil || dll.Tail() != nil {
		t.Errorf("Head and Tail of an empty list should be nil")
	}

	dll.InsertAtEnd(12)                   //insert at the end of an empty list
	dll.InsertAtFront(10)                 //insert at the beginning
	dll.InsertAfter(13, dll.Tail())       //insert after 12 (the tail)
	dll.InsertAfter(11, dll.Head())       //insert after 10 (not the tail)
	dll.InsertBefore(9, dll.Head())       //insert before 10 (the head)
	dll.InsertBefore(14, dll.Tail().Prev) //insert before 12 (not the head)

	checkInt(t, "Head", dll.Head().Data, 9)
	checkInt(t, "Tail", dll.Tail().Data, 13)
	checkBool(t, "Empty", dll.Empty(), false)
	checkInt(t, "Size", dll.Size(), 6)
	checkInts(t, "values", doubleValues(t, dll), []int{9, 10, 11, 14, 12, 13})

	out := captureOutput(dll.Print)
	checkOutput(t, "Print", out, "9 -> 10 -> 11 -> 14 -> 12 -> 13 -> <nil>")

	out = captureOutput(func() { dll.Remove(dll.Head().Next.Next.Next) }) //removing 14 (middle)
	checkOutput(t, "Remove (middle)", out, "removed 14")
	checkInts(t, "values (Remove middle)", doubleValues(t, dll), []int{9, 10, 11, 12, 13})

	out = captureOutput(func() { dll.Remove(dll.Head()) }) //removing 9 (head)
	checkOutput(t, "Remove (head)", out, "removed 9")
	checkInts(t, "values (Remove head)", doubleValues(t, dll), []int{10, 11, 12, 13})

	out = captureOutput(func() { dll.Remove(dll.Tail()) }) //removing 13 (tail)
	checkOutput(t, "Remove (tail)", out, "removed 13")
	checkInts(t, "values (Remove tail)", doubleValues(t, dll), []int{10, 11, 12})

	out = captureOutput(dll.RemoveAtFront) //removing 10
	checkOutput(t, "RemoveAtFront", out, "removed 10")
	out = captureOutput(dll.RemoveAtEnd) //removing 12
	checkOutput(t, "RemoveAtEnd", out, "removed 12")

	checkInt(t, "Head", dll.Head().Data, 11)
	checkInt(t, "Tail", dll.Tail().Data, 11)
	checkInt(t, "Size", dll.Size(), 1)
	checkInts(t, "values", doubleValues(t, dll), []int{11})

	//Remove of the only node empties the list
	captureOutput(func() { dll.Remove(dll.Head()) })
	checkBool(t, "Empty (Remove only node)", dll.Empty(), true)
	if dll.Head() != nil || dll.Tail() != nil {
		t.Errorf("Head and Tail should be nil after removing the only node")
	}

	//RemoveAtFront / RemoveAtEnd of the only node empties the list
	dll.InsertAtEnd(1)
	captureOutput(dll.RemoveAtFront)
	if dll.Head() != nil || dll.Tail() != nil {
		t.Errorf("RemoveAtFront: Head and Tail should be nil after removing the only node")
	}
	dll.InsertAtEnd(1)
	captureOutput(dll.RemoveAtEnd)
	if dll.Head() != nil || dll.Tail() != nil {
		t.Errorf("RemoveAtEnd: Head and Tail should be nil after removing the only node")
	}

	//removing from an empty list does not panic
	out = captureOutput(dll.RemoveAtFront)
	checkOutput(t, "RemoveAtFront (empty)", out, "is empty")
	out = captureOutput(dll.RemoveAtEnd)
	checkOutput(t, "RemoveAtEnd (empty)", out, "is empty")
	checkInt(t, "Size (empty removes)", dll.Size(), 0)

	//InsertAtFront into an empty list sets both Head and Tail
	dll.InsertAtFront(7)
	checkInts(t, "values (InsertAtFront empty)", doubleValues(t, dll), []int{7})
	checkInt(t, "Tail (InsertAtFront empty)", dll.Tail().Data, 7)
	captureOutput(dll.RemoveAtFront)

	//the list is still usable after being emptied
	dll.InsertAtEnd(2)
	dll.InsertAtFront(1)
	dll.InsertAtEnd(3)
	checkInts(t, "values (reuse)", doubleValues(t, dll), []int{1, 2, 3})
	checkInt(t, "Size (reuse)", dll.Size(), 3)

	out = captureOutput(NewDoublyLinkedList[int]().Print)
	checkOutput(t, "Print (empty)", out, "<nil>")
}
