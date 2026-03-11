# class Solution:
#     def isValidSudoku(self, board: List[List[str]]) -> bool:
#         for i in range(9):
#             # row
#             seen = set()
#             for j in range(9):
#                 e = board[i][j]
#                 if e != ".":
#                     if e in seen:
#                         return False
#                     seen.add(e)

#             # col
#             seen = set()
#             for j in range(9):
#                 e = board[j][i]
#                 if e != ".":
#                     if e in seen:
#                         return False
#                     seen.add(e)

#             # box
#             seen = set()
#             boxr = i % 3
#             boxc = i // 3
#             for br in range(3):
#                 for bc in range(3):
#                     e = board[boxr * 3 + br][boxc * 3 + bc]
#                     if e != ".":
#                         if e in seen:
#                             return False
#                         seen.add(e)

#         return True


class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        rows = [set() for _ in range(9)]
        cols = [set() for _ in range(9)]
        boxes = [set() for _ in range(9)]

        for r in range(9):
            for c in range(9):
                e = board[r][c]

                if e == ".":
                    continue

                bidx = (c // 3) + (r // 3) * 3

                if e in rows[r] or e in cols[c] or e in boxes[bidx]:
                    return False

                rows[r].add(e)
                cols[c].add(e)
                boxes[bidx].add(e)

        return True
