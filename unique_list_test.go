package main

import (
	"fmt"
	"testing"
)

func getList(n int, unique bool) *List {
	var counter, data int
	list := new(List)
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
		list := getList(idx, test)
		if list.Unique() != test {
			fmt.Printf("Test #%d failed", idx)
			t.Fail()
		}
	}
}
