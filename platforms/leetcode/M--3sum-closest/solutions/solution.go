import "slices"

func threeSumClosest(nums []int, target int) int {
	slices.Sort(nums)

	closest := nums[0] + nums[1] + nums[2]
	n := len(nums)
	for i := 0; i < n-2; i++ {
		l := i + 1
		r := n - 1
		for l < r {
			s := nums[i] + nums[l] + nums[r]
			if abs(s-target) < abs(closest-target) {
				closest = s
			}

			if s < target {
				l++
			} else if s > target {
				r--
			} else {
				return target
			}
		}

	}
	return closest
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
