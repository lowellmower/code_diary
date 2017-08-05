package diary

import (
	"fmt"
	"testing"

	ds "github.com/lowellmower/code_diary/structures"
)

// GetTestList takes an integer representing how many nodes
// should be created in the list and a boolean representing
// whether or not the returned list should be made unique
func GetTestList(n int, unique bool) *ds.List {
	var counter, data int
	list := new(ds.List)
	for {
		counter++
		if unique {
			data++
		}
		list.Insert(data)
		if counter == n {
			return list
		}
	}
}

func TestUniqueList(t *testing.T) {
	testIO := []bool{
		true,
		true,
		true,
		false,
		true,
		true,
		false,
		false,
	}
	for idx, test := range testIO {
		idx++
		list := GetTestList(idx, test)
		if list.Unique() != test {
			fmt.Printf("Test #%d failed", idx)
			t.Fail()
		}
	}
}
