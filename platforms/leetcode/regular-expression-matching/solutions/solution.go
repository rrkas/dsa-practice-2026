func isMatch(s string, p string) bool {
	sn, pn := len(s), len(p)
	dp := make([][]bool, 0)
	for i := 0; i <= sn; i++ {
		dp = append(dp, make([]bool, pn+1))
	}
	dp[0][0] = true

	// handle patterns like a*, a*b*, etc for empty strings
	for j := 2; j <= pn; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	for i := 1; i <= sn; i++ {
		for j := 1; j <= pn; j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j-2]
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}

	return dp[sn][pn]
}