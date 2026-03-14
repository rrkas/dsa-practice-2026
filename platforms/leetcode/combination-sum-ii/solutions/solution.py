class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates.sort()
        res = []

        def find(start: int, path: list, pathsum: int):
            if pathsum == target:
                res.append(path.copy())

            for i in range(start, len(candidates)):
                if i > start and candidates[i] == candidates[i - 1]:
                    continue

                c = candidates[i]
                if pathsum + c > target:
                    break

                path.append(c)
                find(i + 1, path, pathsum + c)
                path.pop()

        find(0, [], 0)

        return res
