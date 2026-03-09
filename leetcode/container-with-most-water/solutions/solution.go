func maxArea(height []int) int {
	l, r := 0, len(height)-1
	maxArea := 0

	for l < r {
		h := height[l]
		if height[r] < h {
			h = height[r]
		}

		area := h * (r - l)
		if area > maxArea {
			maxArea = area
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxArea
}