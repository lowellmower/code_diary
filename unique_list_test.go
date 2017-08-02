package main

import (
	"fmt"
	"testing"
)

func TestUniqueList(t *testing.T) {
	testIO := []struct {
		list []Nodes
	}{
		{
			[][]int{
				[]int{2, 1},
				[]int{4, 1},
				[]int{2, -3},
				[]int{4, -3},
			},
			[][]int{
				[]int{3, 0},
				[]int{-4, 0},
				[]int{3, -2},
				[]int{-4, -2},
			},
			true,
		},
		{
			[][]int{
				[]int{2, 1},
				[]int{4, 1},
				[]int{2, -3},
				[]int{4, -3},
			},
			[][]int{
				[]int{1, 1},
				[]int{1, 3},
				[]int{-2, 1},
				[]int{-2, 3},
			},
			false,
		},
	}
	for test_idx, test := range testIO {
		result_one := overlappingVector(test.rect_one, test.rect_two)
		result_two := overlappingVector(test.rect_two, test.rect_one)
		if result_one != test.overlap || result_two != test.overlap {
			fmt.Printf("Overlap failed on test %d", test_idx+1)
			t.Fail()
		}
	}
}
