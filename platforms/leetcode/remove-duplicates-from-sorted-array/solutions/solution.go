func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	w := 1

	for r := 1; r < len(nums); r++ {
		if nums[r] != nums[r-1] {
			nums[w] = nums[r]
			w++
		}
	}

	return w
}