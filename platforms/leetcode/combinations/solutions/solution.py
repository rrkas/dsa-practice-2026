class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        res = []

        def find(start, path):
            if len(path) == k:
                res.append(path.copy())
                return

            for i in range(start, n + 1):
                path.append(i)
                find(i + 1, path)
                path.pop()

        find(1, [])
        return res


class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        res = []

        def find(start, path):
            if len(path) == k:
                res.append(path.copy())
                return

            # pruning
            remaining = k - len(path)
            for i in range(start, n - remaining + 2):
                path.append(i)
                find(i + 1, path)
                path.pop()

        find(1, [])
        return res
