func minDistance(word1 string, word2 string) int {
	if word1 == word2 {
		return 0
	}

	m, n := len(word1), len(word2)
	if m == 0 {
		return n
	} else if n == 0 {
		return m
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		row := make([]int, n+1)
		// deletion
		row[0] = i
		dp[i] = row
	}
	// insertion
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] != word2[j-1] {
				dp[i][j] = min(
					dp[i][j-1]+1,   // insert
					dp[i-1][j]+1,   // delete
					dp[i-1][j-1]+1, // replace
				)
			} else {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}

	return dp[m][n]
}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	prev := make([]int, n+1)

	for j := 0; j <= n; j++ {
		prev[j] = j
	}

	for i := 1; i <= m; i++ {
		curr := make([]int, n+1)
		curr[0] = i

		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				curr[j] = prev[j-1]
			} else {
				curr[j] = min(
					curr[j-1]+1,
					prev[j]+1,
					prev[j-1]+1,
				)
			}
		}

		prev = curr
	}

	return prev[n]
}