class Solution:
    def maxSubArray(self, nums):
        curr = nums[0]
        best = nums[0]

        for n in nums[1:]:
            curr = max(n, curr + n)
            best = max(best, curr)

        return best


##################################################


class Solution:
    def maxSubArray(self, nums):
        curr = 0
        best = nums[0]

        for n in nums:
            curr += n
            best = max(best, curr)

            if curr < 0:
                curr = 0

        return best
