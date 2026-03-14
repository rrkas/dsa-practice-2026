class Solution:
    def uniquePathsWithObstacles(self, obstacleGrid: List[List[int]]) -> int:
        if obstacleGrid[0][0] == 1:
            return 0

        m, n = len(obstacleGrid), len(obstacleGrid[0])
        dp = [0] * m
        dp[0] = 1
        for j in range(n):
            for i in range(m):
                if obstacleGrid[i][j] == 1:
                    dp[i] = 0
                elif i > 0:
                    dp[i] += dp[i - 1]
        return dp[m - 1]
