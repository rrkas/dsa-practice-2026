func solveSudoku(board [][]byte) {
	check := func(r, c int, val byte) bool {
		for i := 0; i < 9; i++ {
			br := 3*(r/3) + i/3
			bc := 3*(c/3) + i%3
			if board[r][i] == val || board[i][c] == val || board[br][bc] == val {
				return false
			}
		}
		return true
	}

	var solve func(r, c int) bool
	solve = func(r, c int) bool {
		if r == 9 {
			return true
		}

		nr := r + (c+1)/9
		nc := (c + 1) % 9

		if board[r][c] != '.' {
			return solve(nr, nc)
		}

		for d := byte('1'); d <= '9'; d++ {
			if check(r, c, d) {
				board[r][c] = d
				if solve(nr, nc) {
					return true
				}
				board[r][c] = '.'
			}
		}

		return false
	}

	solve(0, 0)
}