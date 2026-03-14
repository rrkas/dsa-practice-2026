class Solution:
    def combinationSum(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates.sort()
        res = []

        maxdepth = 152

        def find(start: int, path: list, depth: int):
            if depth == maxdepth:
                return

            pathsum = sum(path)

            for i in range(start, len(candidates)):
                c = candidates[i]
                if pathsum + c > target:
                    break

                path.append(c)
                if pathsum + c == target:
                    res.append([*path])

                find(i, path, depth + 1)
                path.pop()

        find(0, [], 0)
        return res
