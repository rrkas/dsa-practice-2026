#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'cavityMap' function below.
#
# The function is expected to return a STRING_ARRAY.
# The function accepts STRING_ARRAY grid as parameter.
#


def cavityMap(grid):
    m, n = len(grid), len(grid[0])
    res = []

    for i in range(m):
        row = list(grid[i])

        for j in range(n):
            if i in [0, m - 1] or j in [0, n - 1]:
                continue

            curr = grid[i][j]

            if (
                curr > grid[i - 1][j]
                and curr > grid[i + 1][j]
                and curr > grid[i][j - 1]
                and curr > grid[i][j + 1]
            ):
                row[j] = "X"

        res.append("".join(row))

    return res


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    n = int(input().strip())
    grid = []

    for _ in range(n):
        grid_item = input()
        grid.append(grid_item)

    result = cavityMap(grid)
    fptr.write("\n".join(result))
    fptr.write("\n")
    fptr.close()
