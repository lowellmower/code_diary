package main

import (
  "fmt"
)

func insertion_sort( s []int ) {
  for idx, _ := range s[1:] {
    for range s[:idx] {
      if (s[idx] < s[idx-1]) {
          tmp := s[idx-1]
          s[idx-1] = s[idx]
          s[idx] = tmp
      }
      idx = idx - 1
    }
  }
}

func main() {
  sort := []int{10, 0, 2, 4, 8, 0, 13, 1, 3, 2}
  fmt.Println(sort)
  insertion_sort(sort)
  fmt.Println(sort)
}
