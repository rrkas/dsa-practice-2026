# class Solution:
#     def solveSudoku(self, board: List[List[str]]) -> None:
#         """
#         Do not return anything, modify board in-place instead.
#         """

#         def check(r, c, val):
#             for i in range(9):
#                 if board[r][i] == val:
#                     return False

#                 if board[i][c] == val:
#                     return False

#                 br = 3 * (r // 3) + i // 3
#                 bc = 3 * (c // 3) + i % 3

#                 if board[br][bc] == val:
#                     return False

#             return True

#         def solve(r, c):
#             if r == 9:
#                 return True

#             nr = r + (c + 1) // 9
#             nc = (c + 1) % 9

#             if board[r][c] != ".":
#                 return solve(nr, nc)

#             for d in "123456789":
#                 if check(r, c, d):
#                     board[r][c] = d
#                     if solve(nr, nc):
#                         return True
#                     board[r][c] = "."

#             return False

#         solve(0, 0)

################################################################


class Solution:
    def solveSudoku(self, board):

        rows = [0] * 9
        cols = [0] * 9
        boxes = [0] * 9
        empty = []

        for r in range(9):
            for c in range(9):
                if board[r][c] == ".":
                    empty.append((r, c))
                else:
                    d = int(board[r][c])
                    mask = 1 << d
                    rows[r] |= mask
                    cols[c] |= mask
                    boxes[(r // 3) * 3 + c // 3] |= mask

        def backtrack(i):
            if i == len(empty):
                return True

            r, c = empty[i]
            b = (r // 3) * 3 + c // 3

            for d in range(1, 10):
                mask = 1 << d

                if rows[r] & mask:
                    continue
                if cols[c] & mask:
                    continue
                if boxes[b] & mask:
                    continue

                board[r][c] = str(d)

                rows[r] |= mask
                cols[c] |= mask
                boxes[b] |= mask

                if backtrack(i + 1):
                    return True

                rows[r] ^= mask
                cols[c] ^= mask
                boxes[b] ^= mask
                board[r][c] = "."

            return False

        backtrack(0)
