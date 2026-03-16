func maxSubArray(nums []int) int {
	curr := nums[0]
	maxs := nums[0]

	for i := 1; i < len(nums); i++ {
		e := nums[i]
		curr = max(e, curr+e)
		maxs = max(maxs, curr)
	}

	return maxs
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

