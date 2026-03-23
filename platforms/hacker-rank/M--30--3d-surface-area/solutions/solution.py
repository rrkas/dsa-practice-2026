#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'surfaceArea' function below.
#
# The function is expected to return an INTEGER.
# The function accepts 2D_INTEGER_ARRAY A as parameter.
#


def surfaceArea(A):
    area = 0
    m, n = len(A), len(A[0])

    # front & back
    area += sum(A[0]) + sum(A[m - 1])

    for i in range(m):
        # left and right
        area += A[i][0] + A[i][n - 1]

        for j in range(n):
            h = A[i][j]
            if h > 0:
                area += 2

            if i > 0:
                area += abs(A[i][j] - A[i - 1][j])

            if j > 0:
                area += abs(A[i][j] - A[i][j - 1])

    return area


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    first_multiple_input = input().rstrip().split()
    H = int(first_multiple_input[0])
    W = int(first_multiple_input[1])
    A = []
    for _ in range(H):
        A.append(list(map(int, input().rstrip().split())))

    result = surfaceArea(A)
    fptr.write(str(result) + "\n")
    fptr.close()
