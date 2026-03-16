class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        m, n = len(matrix), len(matrix[0])

        ridx = -1
        l, r = 0, m - 1
        while l <= r:
            mid = (l + r) // 2
            mv = matrix[mid][0]
            if mv == target:
                return True
            if mv < target:
                l = mid + 1
                ridx = mid
            else:
                r = mid - 1

        if ridx == -1:
            return False

        l, r = 0, n - 1
        while l <= r:
            mid = (l + r) // 2
            mv = matrix[ridx][mid]
            if mv == target:
                return True
            if mv < target:
                l = mid + 1
            else:
                r = mid - 1

        return False


class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        m, n = len(matrix), len(matrix[0])

        l, r = 0, m * n - 1

        while l <= r:
            mid = (l + r) // 2
            val = matrix[mid // n][mid % n]

            if val == target:
                return True
            elif val < target:
                l = mid + 1
            else:
                r = mid - 1

        return False
