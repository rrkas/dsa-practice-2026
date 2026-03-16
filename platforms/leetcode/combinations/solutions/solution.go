func combine(n int, k int) [][]int {
	res := make([][]int, 0)

	var find func(int, []int)
	find = func(start int, path []int) {
		if len(path) == k {
			res = append(res, append([]int{}, path...))
			return
		}

		remaining := k - len(path)
		for i := start; i <= n-remaining+1; i++ {
			path = append(path, i)
			find(i+1, path)
			path = path[:len(path)-1]
		}
	}

	find(1, []int{})

	return res
}

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	path := make([]int, k)

	var dfs func(start, depth int)
	dfs = func(start, depth int) {
		if depth == k {
			comb := make([]int, k)
			copy(comb, path)
			res = append(res, comb)
			return
		}

		remaining := k - depth
		for i := start; i <= n-remaining+1; i++ {
			path[depth] = i
			dfs(i+1, depth+1)
		}
	}

	dfs(1, 0)
	return res
}