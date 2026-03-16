class Solution:
    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        n = len(nums)
        for i in range(n - 1):
            for j in range(i + 1, n):
                if nums[i] > nums[j]:
                    nums[i], nums[j] = nums[j], nums[i]


class Solution:
    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        n = len(nums)
        counts = [0, 0, 0]
        for i in range(n):
            counts[nums[i]] += 1

        i = 0
        for ci, c in enumerate(counts):
            for _ in range(c):
                nums[i] = ci
                i += 1
