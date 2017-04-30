package main

func pairWithSum(list []int, sum int) bool {
	matchSum := make(map[int]bool)

	for _, num := range list {
		ok := matchSum[num]
		if ok {
			return ok
		}
		diff := sum - num
		matchSum[diff] = true
	}

	return false
}
