class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        sign = 1
        if (dividend < 0) != (divisor < 0):
            sign = -1

        a, b = abs(dividend), abs(divisor)
        res = 0
        while a >= b:
            multiple = 1
            while a >= b * (multiple << 1):
                multiple <<= 1

            a -= b * multiple
            res += multiple

        ans = sign * res

        if ans > 2**31 - 1:
            ans = 2**31 - 1
        elif ans < -(2**31):
            ans = -(2**31)

        return ans
