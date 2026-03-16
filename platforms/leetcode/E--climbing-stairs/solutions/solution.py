class Solution:
    def climbStairs(self, n: int) -> int:
        if n < 2:
            return n

        a = 0
        b = 1
        for _ in range(n):
            a, b = b, a + b

        return b
