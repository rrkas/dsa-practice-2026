func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	ridx := -1
	l, r := 0, m-1
	for l <= r {
		mid := (l + r) / 2
		mv := matrix[mid][0]
		if mv == target {
			return true
		}
		if mv > target {
			r = mid - 1
		} else {
			ridx = mid
			l = mid + 1
		}
	}

	if ridx < 0 {
		return false
	}

	l, r = 0, n-1
	for l <= r {
		mid := (l + r) / 2
		mv := matrix[ridx][mid]
		if mv == target {
			return true
		}
		if mv > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return false
}