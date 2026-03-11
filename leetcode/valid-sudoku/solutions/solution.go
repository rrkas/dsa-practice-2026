func isValidSudoku(board [][]byte) bool {
	for i := range 9 {
		// row i
		seen := make(map[byte]int)
		for j := range 9 {
			e := board[i][j]
			if e != '.' {
				seen[e]++
				if seen[e] > 1 {
					// fmt.Println("row failed", j)
					// fmt.Println(seen)
					return false
				}
			}
		}

		// col i
		seen = make(map[byte]int)
		for j := range 9 {
			e := board[j][i]
			if e != '.' {
				seen[e]++
				if seen[e] > 1 {
					// fmt.Println("col failed", j)
					// fmt.Println(seen)
					return false
				}
			}
		}

		// box i
		seen = make(map[byte]int)
		boxr := i % 3
		boxc := i / 3
		for br := range 3 {
			for bc := range 3 {
				e := board[boxr*3+br][boxc*3+bc]
				if e != '.' {
					seen[e]++
					if seen[e] > 1 {
						// fmt.Println("box failed", boxr, boxc)
						// fmt.Println(seen)
						return false
					}
				}
			}
		}
	}
	return true
}