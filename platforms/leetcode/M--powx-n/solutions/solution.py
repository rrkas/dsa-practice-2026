class Solution:
    def myPow(self, x: float, n: int) -> float:
        sign = -1 if n < 0 else 1
        n = abs(n)

        def rec(x, n):
            if n == 0:
                return 1
            if n == 1:
                return x
            if x == 0:
                return 0

            if n % 2 == 0:
                return rec(x * x, n // 2)
            else:
                return x * rec(x, n - 1)

        res = rec(x, n)
        if sign < 0:
            res = 1 / res
        return res
