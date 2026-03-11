func searchRange(nums []int, target int) []int {
	n := len(nums)
	l, r := 0, n-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			start, end := mid, mid
			for start > 0 && nums[start-1] == target {
				start--
			}
			for end < n-1 && nums[end+1] == target {
				end++
			}
			return []int{start, end}
		}

		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return []int{-1, -1}
}