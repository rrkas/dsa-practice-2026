func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	n := len(nums)

	var find func(int, []int)
	find = func(start int, path []int) {
		res = append(res, append([]int{}, path...))

		for i := start; i < n; i++ {
			path = append(path, nums[i])
			find(i+1, path)
			path = path[:len(path)-1]
		}
	}
	find(0, []int{})

	return res
}