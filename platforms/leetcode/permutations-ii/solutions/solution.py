class Solution:
    def permuteUnique(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        n = len(nums)
        seen = [False] * n
        res = []

        def perm(path):
            if len(path) == n:
                res.append(path[:])
                return

            i = 0
            while i < n:
                if seen[i]:
                    i += 1
                    continue

                if i > 0 and nums[i] == nums[i - 1] and seen[i - 1]:
                    i += 1
                    continue

                path.append(nums[i])
                seen[i] = True
                perm(path)
                path.pop()
                seen[i] = False
                i += 1

        perm([])
        return res
