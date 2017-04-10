package main

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	testIO := []struct{
		sort []int
		exp  []int
	}{
		{
			[]int{5,4,3,2,1,0},
			[]int{0,1,2,3,4,5},
		},
		{
			[]int{5,5,5,5,5,0},
			[]int{0,5,5,5,5,5},
		},
		{
			[]int{0,0,1},
			[]int{0,0,1},
		},
		{
			[]int{1,0,0},
			[]int{0,0,1},
		},
		{
			[]int{0,1,0},
			[]int{0,0,1},
		},
		{
			[]int{1,-2,-3,4,0},
			[]int{-3,-2,0,1,4},
		},
		{
			[]int{0,0,0,0,1,0,0,0,0,0},
			[]int{0,0,0,0,0,0,0,0,0,1},
		},
	}

	for _, test := range testIO {
		insertionSort(test.sort)
		for i, n := range test.sort {
			if n != test.exp[i] {
				fmt.Printf("n: %d  e: %d\n", n, test.exp[i])
				t.Fail()
			}
		}
	}
}
