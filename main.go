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

package main

import (
	"fmt"
	"github.com/lowellmower/code_diary/diary"
)

func main() {
	list := diary.NewList()
	fmt.Println(list)
}
