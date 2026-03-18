func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])

	var find func(int, int, int) bool
	find = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}

		if r < 0 || c < 0 || r >= m || c >= n || board[r][c] != word[i] {
			return false
		}

		temp := board[r][c]
		board[r][c] = '#' // mark visited

		found := (find(r-1, c, i+1) ||
			find(r+1, c, i+1) ||
			find(r, c-1, i+1) ||
			find(r, c+1, i+1))

		board[r][c] = temp
		return found
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if find(i, j, 0) {
				return true
			}
		}
	}

	return false
}