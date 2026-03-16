func maxArea(height []int) int {
	l, r, maxA := 0, len(height)-1, 0

	for l < r {
		if h := min(height[l], height[r]) * (r - l); h > maxA {
			maxA = h
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxA
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}