func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])

	res := make([]int, 0)

	top, bottom, left, right := 0, m-1, 0, n-1
	var i, j int
	for top <= bottom && left <= right {
		for i = left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++
		for j = top; j <= bottom; j++ {
			res = append(res, matrix[j][right])
		}
		right--
		if top <= bottom {
			for i = right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}
		if left <= right {
			for j = bottom; j >= top; j-- {
				res = append(res, matrix[j][left])
			}
			left++
		}
	}

	return res
}