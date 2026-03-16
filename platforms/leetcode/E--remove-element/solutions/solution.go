func removeElement(nums []int, val int) int {
	w := 0
	for r := range nums {
		if nums[r] != val {
			nums[w] = nums[r]
			w++
		}
	}
	return w
}