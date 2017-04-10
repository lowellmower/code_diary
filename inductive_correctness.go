package main

import (
	"fmt"
)

func increment(n int) int {
	if n == 0 {
		return 1
	}
	if n % 2 == 1 {
		return 2 * increment(n/2)
	} else {
		return n + 1
	}
}

func main() {
	n := increment(10)
	fmt.Println(n)
}
