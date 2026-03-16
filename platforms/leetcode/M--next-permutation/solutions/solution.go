func nextPermutation(nums []int) {
	n := len(nums)

	// 1. find pivot
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// 2. find next larger element
	if i >= 0 {
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 3. reverse suffix
	l, r := i+1, n-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}