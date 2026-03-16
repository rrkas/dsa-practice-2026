class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        w = 1
        for r in range(1, len(nums)):
            if nums[r] != nums[r - 1]:
                nums[w] = nums[r]
                w += 1
        return w
