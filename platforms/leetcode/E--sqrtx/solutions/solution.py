class Solution:
    def mySqrt(self, x: int) -> int:
        if x < 2:
            return x

        l, r = 0, x
        ans = 0

        while l <= r:
            mid = l + (r - l) // 2

            if mid <= x // mid:
                ans = mid
                l = mid + 1
            else:
                r = mid - 1

        return ans
