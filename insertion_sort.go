package main

func insertionSort(s []int) {
	for i, _ := range s {
		if i > 0 {
			for range s[:i] {
				if s[i] < s[i-1] {
					tmp := s[i-1]
					s[i-1] = s[i]
					s[i] = tmp
				}
				i = i - 1
			}
		}
	}
}
