import "slices"

func combinationSum(candidates []int, target int) [][]int {
	slices.Sort(candidates)

	maxdepth := 152
	res := make([][]int, 0)

	var find func(int, []int, int)
	find = func(start int, path []int, depth int) {
		if depth == maxdepth {
			return
		}

		pathsum := 0
		for _, step := range path {
			pathsum += step
		}

		for i := start; i < len(candidates); i++ {
			c := candidates[i]
			if pathsum+c > target {
				break
			}
			path = append(path, c)
			if pathsum+c == target {
				copyPath := append([]int{}, path...)
				res = append(res, copyPath)
			}
			find(i, path, depth+1)
			path = path[:len(path)-1]
		}
	}

	find(0, make([]int, 0), 0)

	return res
}