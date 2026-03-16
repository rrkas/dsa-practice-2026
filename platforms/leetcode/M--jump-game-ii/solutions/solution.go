func jump(nums []int) int {
	jumps := 0
	curr_end := 0
	far := 0
	for i, e := range nums {
		if i == len(nums)-1 {
			break
		}

		if i+e > far {
			far = i + e
		}

		if i == curr_end {
			jumps++
			curr_end = far
		}
	}
	return jumps
}