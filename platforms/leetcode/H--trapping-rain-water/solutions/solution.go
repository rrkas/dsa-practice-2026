func trap(height []int) int {
	n := len(height)
	l, r := 0, n-1
	lmax, rmax := 0, 0
	water := 0

	for l < r {
		if height[l] < height[r] {
			if height[l] >= lmax {
				lmax = height[l]
			} else {
				water += lmax - height[l]
			}
			l++
		} else {
			if height[r] >= rmax {
				rmax = height[r]
			} else {
				water += rmax - height[r]
			}
			r--
		}
	}
	return water
}