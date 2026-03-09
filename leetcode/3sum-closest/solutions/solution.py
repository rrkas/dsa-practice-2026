class Solution:
    def threeSumClosest(self, nums: List[int], target: int) -> int:
        nums.sort()

        closest = sum(nums[:3])
        for i in range(len(nums) - 2):
            l = i + 1
            r = len(nums) - 1

            while l < r:
                s = nums[i] + nums[l] + nums[r]
                if abs(target - s) < abs(target - closest):
                    closest = s

                if s < target:
                    l += 1
                elif s > target:
                    r -= 1
                else:
                    return target

        return closest
