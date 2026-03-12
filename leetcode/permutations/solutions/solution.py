class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        res = []
        n = len(nums)
        used = [False] * n

        def perm(path):
            if len(path) == n:
                res.append(path.copy())

            for i, e in enumerate(nums):
                if used[i]:
                    continue

                path.append(e)
                used[i] = True
                perm(path)
                path.pop()
                used[i] = False

        perm([])

        return res
