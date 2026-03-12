func canJump(nums []int) bool {
	curr, far, n := 0, 0, len(nums)

	for i := 0; i < n; i++ {
		far = min(n-1, max(far, i+nums[i]))

		if i == far {
			break
		}
		if i == curr {
			curr = far
		}
	}

	return far == n-1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
