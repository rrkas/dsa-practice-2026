import "slices"

func subsetsWithDup(nums []int) [][]int {
	slices.Sort(nums)

	res := make([][]int, 0)

	var find func(int, []int)
	find = func(start int, subset []int) {
		res = append(res, append([]int{}, subset...))

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			subset = append(subset, nums[i])
			find(i+1, subset)
			subset = subset[:len(subset)-1]
		}
	}

	find(0, []int{})

	return res
}