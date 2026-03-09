class Solution:
    def maxArea(self, height: List[int]) -> int:
        maxa = 0
        l, r = 0, len(height) - 1
        while l < r:
            hl = height[l]
            hr = height[r]
            area = min(hl, hr) * (r - l)
            maxa = max(maxa, area)
            if hl < hr:
                l += 1
            else:
                r -= 1
        return maxa
