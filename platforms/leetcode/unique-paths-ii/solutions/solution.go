func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	if obstacleGrid[0][0] == 1 {
		return 0
	}

	dp[0][0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}

			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}

			if j > 0 {
				dp[i][j] += dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

func uniquePathsWithObstacles(grid [][]int) int {
    m, n := len(grid), len(grid[0])

    dp := make([]int, n)

    if grid[0][0] == 1 {
        return 0
    }

    dp[0] = 1

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {

            if grid[i][j] == 1 {
                dp[j] = 0
            } else if j > 0 {
                dp[j] += dp[j-1]
            }

        }
    }

    return dp[n-1]
}