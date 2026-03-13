class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        n, m = len(s), len(p)

        dp = [[False] * (m + 1) for _ in range(n + 1)]

        dp[0][0] = True

        for j in range(1, m + 1):
            if p[j - 1] == "*":
                dp[0][j] = dp[0][j - 1]

        for i in range(1, n + 1):
            for j in range(1, m + 1):

                if p[j - 1] == "?" or s[i - 1] == p[j - 1]:
                    dp[i][j] = dp[i - 1][j - 1]

                elif p[j - 1] == "*":
                    dp[i][j] = dp[i][j - 1] or dp[i - 1][j]

        return dp[n][m]


class Solution:
    def isMatch(self, s: str, p: str) -> bool:

        i = 0
        j = 0
        star = -1
        match = 0

        while i < len(s):

            if j < len(p) and (p[j] == s[i] or p[j] == "?"):
                i += 1
                j += 1

            elif j < len(p) and p[j] == "*":
                star = j
                match = i
                j += 1

            elif star != -1:
                j = star + 1
                match += 1
                i = match

            else:
                return False

        while j < len(p) and p[j] == "*":
            j += 1

        return j == len(p)
