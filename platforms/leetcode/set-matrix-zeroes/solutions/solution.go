func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	zrows, zcols := make([]bool, m), make([]bool, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				zrows[i] = true
				zcols[j] = true
			}
		}
	}

	for i := 0; i < m; i++ {
		if zrows[i] {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		} else {
			for j := 0; j < n; j++ {
				if zcols[j] {
					matrix[i][j] = 0
				}
			}
		}
	}
}