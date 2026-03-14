func generateMatrix(n int) [][]int {
	c := 1
	top, bottom := 0, n-1
	left, right := 0, n-1

	matrix := make([][]int, 0)
	for t := 0; t < n; t++ {
		matrix = append(matrix, make([]int, n))
	}

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			matrix[top][i] = c
			c++
		}
		top++
		for j := top; j <= bottom; j++ {
			matrix[j][right] = c
			c++
		}
		right--
		for i := right; i >= left; i-- {
			matrix[bottom][i] = c
			c++
		}
		bottom--
		for j := bottom; j >= top; j-- {
			matrix[j][left] = c
			c++
		}
		left++
	}

	return matrix
}