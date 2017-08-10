// This specific problem was introduced to me on August 6th 2017 over coffee with
// a co-worker. I enjoy this problem, because, as he pointed out it is not overly
// difficult but is deceptively tricky. We talked through it and I pseudo coded it
// on a piece of paper and am now legitimately implementing the solution here.

// Problem:
// Given a linked list where the nodes only know about their respective next,
// write a function remove(d) which will remove the node from the list.

// Example:
// l = ['x', 'n', 'm', 'v', 'n']
// l.remove('n')
// l = ['x', 'm', 'v']
package diary

import (
	// "fmt"
	"testing"

	ds "github.com/lowellmower/code_diary/structures"
)

func TestRemoveNode(t *testing.T) {
	testIO := []struct {
		list *ds.List
		exp  int
	}{
		{
			list: GetTestList(2, true),
			exp:  1,
		},
		{
			list: GetTestList(3, true),
			exp:  2,
		},
		{
			list: GetTestList(4, true),
			exp:  3,
		},
	}

	for idx, test := range testIO {
		ogLen := test.list.Length()
		if ogLen == test.exp {
			t.Fatalf("Length() returned incorrectly: %d", ogLen)
		}
		idx++
		test.list.Remove(idx)
		if test.list.Length() == ogLen {
			t.Fatal("Remove() failed, length equal to original")
		}
		if test.list.Length() != test.exp {
			t.Fatal("Remove() failed to remove all instances of data")
		}

		// More obvious "manual" testing
		l := new(ds.List)
		if l.Length() != 0 {
			t.Fatalf("LIST: %#v", l)
		}

		l.Insert(1)
		if l.Length() != 1 {
			t.Fatalf("LIST: %#v", l)
		}

		// Remove the only thing in the list
		l.Remove(1)
		if l.Length() != 0 {
			t.Fatalf("LIST: %#v", l)
		}

		// Make the same list larger
		l.Insert("a")
		l.Insert("b")
		l.Insert("c")
		l.Insert("d")
		if l.Length() != 4 {
			t.Fatalf("LIST: %#v", l)
		}

		// Remove something from the middle
		l.Remove("c")
		if l.Length() != 3 {
			t.Fatalf("LIST: %#v", l)
		}

		// Insert (append) a different type
		l.Insert(3)
		if l.Length() != 4 {
			t.Fatalf("LIST: %#v", l)
		}

		// Remove the last thing and ensure only strings remain
		l.Remove(3)
		if l.Length() != 3 {
			t.Fatalf("LIST: %#v", l)
		}
		if l.Head.Data != "a" {
			t.Fatalf("LIST: %#v", l)
		}
		if l.Head.Next.Data != "b" {
			t.Fatalf("LIST: %#v", l)
		}
		if l.Head.Next.Next.Data != "d" && l.Head.Next.Next.Next == nil {
			t.Logf("DATA: %d", l.Head.Next.Next.Data)
			t.Fatalf("LIST: %#v", l.Head.Next.Next.Next)
		}

		// make a new list of non unique types and remove all
		identicalList := GetTestList(2, false)
		if identicalList.Length() != 2 {
			t.Fatalf("LIST: %#v", identicalList)
		}
		if identicalList.Head.Data != identicalList.Head.Next.Data {
			t.Logf("NEXT: %#v", identicalList.Head.Next.Data)
			t.Logf("HEAD: %#v", identicalList.Head.Data)
			t.Fatalf("LIST: %#v", identicalList)
		}
		identicalList.Remove(0)
		if !identicalList.IsEmpty() {
			t.Fatalf("HEAD: %#v", identicalList)
		}

		identicalList = GetTestList(3, false)
		if identicalList.Length() != 3 {
			t.Fatalf("LIST: %#v", identicalList)
		}
		if identicalList.Head.Data != identicalList.Head.Next.Data {
			t.Logf("NEXT: %#v", identicalList.Head.Next.Data)
			t.Logf("HEAD: %#v", identicalList.Head.Data)
			t.Fatalf("LIST: %#v", identicalList)
		}
		identicalList.Remove(0)
		if !identicalList.IsEmpty() {
			t.Logf("NEXT: %#v", identicalList.Head.Next)
			t.Logf("HEAD: %#v", identicalList.Head)
			t.Fatalf("LIST: %#v", identicalList)
		}

	}
}
