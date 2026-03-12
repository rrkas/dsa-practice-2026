class Solution:
    def solveNQueens(self, n: int) -> List[List[str]]:
        res = []

        Q = "Q"
        DOT = "."

        board = [[DOT] * n for _ in range(n)]

        def check(r, c):
            # columns
            for ri in range(n):
                if board[ri][c] == Q:
                    return False

            i, j = r - 1, c - 1
            while i >= 0 and j >= 0:
                if board[i][j] == Q:
                    return False
                i -= 1
                j -= 1

            i, j = r - 1, c + 1
            while i >= 0 and j < n:
                if board[i][j] == Q:
                    return False
                i -= 1
                j += 1

            return True

        def find(r):
            if r == n:
                res.append(["".join(row) for row in board])
                return

            for c in range(n):
                if check(r, c):
                    board[r][c] = Q
                    find(r + 1)
                    board[r][c] = DOT

        find(0)

        return res
