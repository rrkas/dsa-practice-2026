#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'bomberMan' function below.
#
# The function is expected to return a STRING_ARRAY.
# The function accepts following parameters:
#  1. INTEGER n
#  2. STRING_ARRAY grid
#


def fullgrid(m, n):
    return ["O" * n] * m


def detonate(grid, m, n):
    newGrid = [[False] * n for _ in range(m)]

    for i in range(m):
        for j in range(n):
            if grid[i][j]:
                continue

            valid = True

            if valid and i > 0 and grid[i - 1][j]:
                valid = False

            if valid and i + 1 < m and grid[i + 1][j]:
                valid = False

            if valid and j > 0 and grid[i][j - 1]:
                valid = False

            if valid and j + 1 < n and grid[i][j + 1]:
                valid = False

            newGrid[i][j] = valid

    return newGrid


def gridstr(grid, m, n):
    res = []
    for i in range(m):
        s = "".join(["O" if grid[i][j] else "." for j in range(n)])
        res.append(s)
    return res


def bomberMan(t, grid, m, n):
    if t == 1:
        return grid

    if t in [2, 4]:
        return fullgrid(m, n)

    igrid = [[c == "O" for c in e] for e in grid]

    A = detonate(igrid, m, n)
    if t == 3:
        return gridstr(A, m, n)

    B = detonate(A, m, n)
    if t == 5:
        return gridstr(B, m, n)

    t -= 6
    d = t % 4
    match d:
        case 0:
            return fullgrid(m, n)
        case 1:
            return gridstr(A, m, n)
        case 2:
            return fullgrid(m, n)
        case 3:
            return gridstr(B, m, n)


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    first_multiple_input = input().rstrip().split()
    r = int(first_multiple_input[0])
    c = int(first_multiple_input[1])
    n = int(first_multiple_input[2])
    grid = []

    for _ in range(r):
        grid_item = input()
        grid.append(grid_item)

    result = bomberMan(n, grid, r, c)
    fptr.write("\n".join(result))
    fptr.write("\n")
    fptr.close()
