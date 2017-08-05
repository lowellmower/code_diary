package structures

import (
	"testing"
)

func testList(n int) *List {
	var counter int
	list := new(List)
	for {
		counter++
		list.Insert(counter)
		if counter == n {
			return list
		}
	}
}

func TestLength(t *testing.T) {
	testIO := []struct {
		list *List
		exp  int
	}{
		{
			list: new(List),
			exp:  0,
		},
		{
			list: testList(1),
			exp:  1,
		},
		{
			list: testList(2),
			exp:  2,
		},
		{
			list: testList(3129),
			exp:  3129,
		},
	}

	for _, test := range testIO {
		if test.list.Length() != test.exp {
			t.Fail()
		}
	}
}
