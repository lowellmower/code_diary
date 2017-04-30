package main

func max(vectors [][]int) (max_x int, max_y int) {
	max_x = vectors[0][0]
	max_y = vectors[1][1]
	for _, v := range vectors {
		if v[0] > max_x {
			max_x = v[0]
		}
		if v[1] > max_y {
			max_y = v[1]
		}
	}
	return
}

func min(vectors [][]int) (min_x int, min_y int) {
	min_x = vectors[0][0]
	min_y = vectors[1][1]
	for _, v := range vectors {
		if v[0] < min_x {
			min_x = v[0]
		}
		if v[1] < min_y {
			min_y = v[1]
		}
	}
	return
}

// overlappingVector will return a boolean value representing
// the truthiness of whether two rectangles represented by
// multidimensional arrays of vector points overlap
func overlappingVector(r1 [][]int, r2 [][]int) bool {
	for _, v := range r1 {
		r2_mx_x, r2_mx_y := max(r2)
		r2_mn_x, r2_mn_y := min(r2)
		if v[0] < r2_mx_x && v[0] > r2_mn_x {
			return true
		}
		if v[1] < r2_mx_y && v[1] > r2_mn_y {
			return true
		}

	}
	return false
}
