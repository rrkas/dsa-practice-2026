import "slices"

func fourSum(nums []int, target int) [][]int {
	slices.Sort(nums)
	res := make([][]int, 0)
	n := len(nums)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			l := j + 1
			r := n - 1
			for l < r {
				s := nums[i] + nums[j] + nums[l] + nums[r]
				if s == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					r--

					for l < r && nums[l] == nums[l-1] {
						l++
					}

					for l < r && nums[r] == nums[r+1] {
						r--
					}
				} else if s < target {
					l++
				} else {
					r--
				}
			}
		}
	}
	return res
}