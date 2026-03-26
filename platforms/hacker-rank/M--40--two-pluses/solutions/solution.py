#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'twoPluses' function below.
#
# The function is expected to return an INTEGER.
# The function accepts STRING_ARRAY grid as parameter.
#


def twoPluses(grid):
    m, n = len(grid), len(grid[0])

    u, d, l, r = [], [], [], []
    for i in range(m):
        u.append([0] * n)  # v: u to d
        d.append([0] * n)  # v: d to u
        l.append([0] * n)  # h: l to r
        r.append([0] * n)  # h: r to l

    for i in range(m):
        for j in range(n):
            if grid[i][j] == "G":
                u[i][j] = 1
                l[i][j] = 1

                if i > 0:
                    u[i][j] += u[i - 1][j]

                if j > 0:
                    l[i][j] += l[i][j - 1]

    for i in range(m - 1, -1, -1):
        for j in range(n - 1, -1, -1):
            if grid[i][j] == "G":
                d[i][j] = 1
                r[i][j] = 1

                if i + 1 < m:
                    d[i][j] += d[i + 1][j]

                if j + 1 < n:
                    r[i][j] += r[i][j + 1]

    def cidx(i, j):
        return i * n + j

    pluses = []
    for i in range(m):
        for j in range(n):
            if grid[i][j] != "G":
                continue

            maxarm = min(u[i][j], d[i][j], l[i][j], r[i][j]) - 1

            for k in range(maxarm + 1):
                mask = 0
                mask |= 1 << cidx(i, j)

                for d_ in range(k + 1):
                    mask |= 1 << cidx(i - d_, j)
                    mask |= 1 << cidx(i + d_, j)
                    mask |= 1 << cidx(i, j - d_)
                    mask |= 1 << cidx(i, j + d_)

                area = 1 + 4 * k
                pluses.append((area, mask))

    pluses.sort(reverse=True)

    maxprod = 0

    for i in range(len(pluses) - 1):
        for j in range(i + 1, len(pluses)):
            prod = pluses[i][0] * pluses[j][0]
            if prod <= maxprod:
                break

            if pluses[i][1] & pluses[j][1] == 0:
                maxprod = prod

    return maxprod


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    first_multiple_input = input().rstrip().split()
    n = int(first_multiple_input[0])
    m = int(first_multiple_input[1])
    grid = []

    for _ in range(n):
        grid_item = input()
        grid.append(grid_item)

    result = twoPluses(grid)
    fptr.write(str(result) + "\n")
    fptr.close()
