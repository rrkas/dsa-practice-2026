import "strings"

func solveNQueens(n int) [][]string {
	// . = false
	// Q = true
	board := make([][]bool, 0, n)
	for i := 0; i < n; i++ {
		row := make([]bool, n, n)
		board = append(board, row)
	}

	res := make([][]string, 0)

	cols := make([]bool, n)
	diag := make([]bool, (2 * n))
	altdiag := make([]bool, (2 * n))

	check := func(r, c int) bool {
		if cols[c] {
			return false
		}

		if diag[r+c] {
			return false
		}

		if altdiag[r-c+n] {
			return false
		}

		return true
	}

	toggle := func(r, c int) {
		board[r][c] = !board[r][c]
		cols[c] = !cols[c]
		diag[r+c] = !diag[r+c]
		altdiag[r-c+n] = !altdiag[r-c+n]
	}

	var find func(int)
	find = func(r int) {
		if r == n {
			bcopy := make([]string, 0, n)
			for i := 0; i < n; i++ {
				sb := strings.Builder{}
				for j := 0; j < n; j++ {
					if board[i][j] {
						sb.WriteString("Q")
					} else {
						sb.WriteString(".")
					}
				}
				bcopy = append(bcopy, sb.String())
			}
			res = append(res, bcopy)
		}

		for c := 0; c < n; c++ {
			if check(r, c) {
				toggle(r, c)
				find(r + 1)
				toggle(r, c)
			}
		}

	}

	find(0)
	return res
}

