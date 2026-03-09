# class Solution:
#     def reverse(self, x: int) -> int:
#         mul = -1 if x < 0 else 1
#         x *= mul
#         x = mul*int(str(x)[::-1])

#         val2e31 =2**31
#         if x < -val2e31 or x > val2e31-1:
#             return 0

#         return x


class Solution:
    def reverse(self, x: int) -> int:
        res = 0
        sign = -1 if x < 0 else 1
        x *= sign

        while x:
            digit = x % 10
            x //= 10

            if res > (2**31 - 1) // 10:
                return 0

            res = res * 10 + digit

        return sign * res
