func sortColors(nums []int) {
	n := len(nums)
	counts := []int{0, 0, 0}
	for i := 0; i < n; i++ {
		counts[nums[i]]++
	}
	i := 0
	for ci := 0; ci < 3; ci++ {
		for j := 0; j < counts[ci]; j++ {
			nums[i] = ci
			i++
		}
	}
}