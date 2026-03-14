class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        n = len(nums)
        l, r = 0, n - 1
        while l <= r:
            mid = (l + r) // 2

            if nums[mid] == target:
                start, end = mid, mid
                while start > 0 and nums[start - 1] == target:
                    start -= 1
                while end < n - 1 and nums[end + 1] == target:
                    end += 1
                return [start, end]

            if nums[mid] < target:
                l = mid + 1
            else:
                r = mid - 1
        return [-1, -1]
