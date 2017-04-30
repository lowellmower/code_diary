package main

import "fmt"

func main() {
	err := fmt.Errorf("We've got %d error", 1)

	if err != nil {
		fmt.Println(err)
	}
}
