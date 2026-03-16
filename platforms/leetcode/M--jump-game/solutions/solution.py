class Solution:
    def canJump(self, nums: List[int]) -> bool:
        far = 0
        curr = 0
        n = len(nums)

        for i in range(len(nums)):
            e = nums[i]
            far = min(n - 1, max(far, i + e))

            if far == i:
                break

            if i == curr:
                curr = far

        return far == len(nums) - 1
