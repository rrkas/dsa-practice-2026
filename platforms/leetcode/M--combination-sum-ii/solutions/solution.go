import "slices"

func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	res := make([][]int, 0)

	var find func(int, []int, int)

	find = func(start int, path []int, pathsum int) {
		if pathsum == target {
			res = append(res, append([]int{}, path...))
		}

		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			c := candidates[i]
			if pathsum+c > target {
				break
			}

			path = append(path, c)
			find(i+1, path, pathsum+c)
			path = path[:len(path)-1]
		}
	}

	find(0, []int{}, 0)

	return res
}