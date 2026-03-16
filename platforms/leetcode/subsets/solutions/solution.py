class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        res = []
        n = len(nums)

        def find(start, path):
            res.append(path[:])

            for i in range(start, n):
                path.append(nums[i])
                find(i + 1, path)
                path.pop()

        find(0, [])

        return res
