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
		// {
		// 	list: new(ds.List),
		// 	exp:  0,
		// },
		{
			list: GetTestList(2, true),
			exp:  1,
		},
		// {
		// 	list: GetTestList(3, true),
		// 	exp:  2,
		// },
	}

	for _, test := range testIO {
		ogLen := test.list.Length()
		if ogLen == test.exp {
			t.Fail()
		}
	}
}
