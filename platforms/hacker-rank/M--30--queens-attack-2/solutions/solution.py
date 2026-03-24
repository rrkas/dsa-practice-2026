#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'queensAttack' function below.
#
# The function is expected to return an INTEGER.
# The function accepts following parameters:
#  1. INTEGER n
#  2. INTEGER k
#  3. INTEGER r_q
#  4. INTEGER c_q
#  5. 2D_INTEGER_ARRAY obstacles
#


# def queensAttack(n, k, r_q, c_q, obstacles):
#     def crc(n, r, c):
#         return n - r, c - 1

#     def checkcoord(coord):
#         for e in obstacles:
#             if e[0] == coord[0] and e[1] == coord[1]:
#                 return False
#         return True

#     r_q, c_q = crc(n, r_q, c_q)
#     obstacles = [crc(n, *e) for e in obstacles]
#     res = 0

#     # vertical
#     for i in range(r_q - 1, -1, -1):
#         if checkcoord([i, c_q]):
#             res += 1
#         else:
#             break

#     for i in range(r_q + 1, n):
#         if checkcoord([i, c_q]):
#             res += 1
#         else:
#             break

#     # horizontal
#     for i in range(c_q - 1, -1, -1):
#         if checkcoord([r_q, i]):
#             res += 1
#         else:
#             break

#     for i in range(c_q + 1, n):
#         if checkcoord([r_q, i]):
#             res += 1
#         else:
#             break

#     # diagonals
#     i, j = r_q - 1, c_q - 1
#     while i >= 0 and j >= 0:
#         if checkcoord([i, j]):
#             res += 1
#         else:
#             break
#         i -= 1
#         j -= 1

#     i, j = r_q + 1, c_q + 1
#     while i < n and j < n:
#         if checkcoord([i, j]):
#             res += 1
#         else:
#             break
#         i += 1
#         j += 1

#     i, j = r_q - 1, c_q + 1
#     while i >= 0 and j < n:
#         if checkcoord([i, j]):
#             res += 1
#         else:
#             break
#         i -= 1
#         j += 1

#     i, j = r_q + 1, c_q - 1
#     while i < n and j >= 0:
#         if checkcoord([i, j]):
#             res += 1
#         else:
#             break
#         i += 1
#         j -= 1

#     return res


def queensAttack(n, k, r_q, c_q, obstacles):
    # initial limits (board edges)
    up = n - r_q
    down = r_q - 1
    right = n - c_q
    left = c_q - 1

    up_right = min(up, right)
    up_left = min(up, left)
    down_right = min(down, right)
    down_left = min(down, left)

    for r, c in obstacles:
        # same column
        if c == c_q:
            if r > r_q:
                up = min(up, r - r_q - 1)
            else:
                down = min(down, r_q - r - 1)

        # same row
        elif r == r_q:
            if c > c_q:
                right = min(right, c - c_q - 1)
            else:
                left = min(left, c_q - c - 1)

        # diagonals
        elif abs(r - r_q) == abs(c - c_q):
            if r > r_q and c > c_q:
                up_right = min(up_right, r - r_q - 1)
            elif r > r_q and c < c_q:
                up_left = min(up_left, r - r_q - 1)
            elif r < r_q and c > c_q:
                down_right = min(down_right, r_q - r - 1)
            elif r < r_q and c < c_q:
                down_left = min(down_left, r_q - r - 1)

    return up + down + left + right + up_right + up_left + down_right + down_left


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    first_multiple_input = input().rstrip().split()
    n = int(first_multiple_input[0])
    k = int(first_multiple_input[1])
    second_multiple_input = input().rstrip().split()
    r_q = int(second_multiple_input[0])
    c_q = int(second_multiple_input[1])
    obstacles = []
    for _ in range(k):
        obstacles.append(list(map(int, input().rstrip().split())))

    result = queensAttack(n, k, r_q, c_q, obstacles)
    fptr.write(str(result) + "\n")
    fptr.close()
