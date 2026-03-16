class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        zrows = 0
        zcols = 0

        m, n = len(matrix), len(matrix[0])

        for i in range(m):
            for j in range(n):
                if matrix[i][j] == 0:
                    zrows |= 1 << i
                    zcols |= 1 << j

        for i in range(m):
            for j in range(n):
                if (zrows >> i) & 1 or (zcols >> j) & 1:
                    matrix[i][j] = 0


class Solution:
    def setZeroes(self, matrix: List[List[int]]) -> None:
        m, n = len(matrix), len(matrix[0])

        zrows = 0
        zcols = 0

        for i in range(m):
            for j in range(n):
                if matrix[i][j] == 0:
                    zrows |= 1 << i
                    zcols |= 1 << j

        for i in range(m):
            if (zrows >> i) & 1:
                matrix[i] = [0] * n
            else:
                for j in range(n):
                    if (zcols >> j) & 1:
                        matrix[i][j] = 0
