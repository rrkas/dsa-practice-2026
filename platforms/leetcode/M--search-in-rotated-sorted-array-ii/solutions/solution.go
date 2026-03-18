func search(nums []int, target int) bool {
	n := len(nums)
	l, r := 0, n-1

	for l <= r {
		mid := (l + r) / 2

		if nums[mid] == target {
			return true
		}

		if nums[l] <= nums[mid] {
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[r] >= target && target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return false
}