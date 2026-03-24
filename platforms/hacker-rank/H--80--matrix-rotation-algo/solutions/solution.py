#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'matrixRotation' function below.
#
# The function accepts following parameters:
#  1. 2D_INTEGER_ARRAY matrix
#  2. INTEGER r
#


def allSameCheck(matrix, m, n):
    e = matrix[0][0]
    for i in range(m):
        for j in range(n):
            if e != matrix[i][j]:
                return False
    return True


### Complexity = O(turns * perimeter)
# def rotateLayer(matrix, turns, l, r, t, b):
#     peri = 2 * ((b - t) + (r - l))
#     turns = turns % peri

#     if turns == 0:
#         return

#     for _ in range(turns):
#         # top left element (corner element)
#         tle = matrix[t][l]

#         # shift top row to left
#         for j in range(l, r):
#             matrix[t][j] = matrix[t][j + 1]

#         # shift right col to up
#         for i in range(t, b):
#             matrix[i][r] = matrix[i + 1][r]

#         # shift bottom row to right
#         for j in range(r, l, -1):
#             matrix[b][j] = matrix[b][j - 1]

#         # shift left col to down
#         for i in range(b, t, -1):
#             matrix[i][l] = matrix[i - 1][l]

#         # top left element place in new position (corner element)
#         matrix[t + 1][l] = tle


### CHATGPT SOLN
### Complexity = O(perimeter)
def rotateLayer(matrix, turns, l, r, t, b):
    peri = 2 * ((b - t) + (r - l))
    turns = turns % peri

    if turns == 0:
        return

    # Step 1: extract the layer into a 1D array
    arr = []

    # top row
    for j in range(l, r + 1):
        arr.append(matrix[t][j])

    # right column
    for i in range(t + 1, b):
        arr.append(matrix[i][r])

    # bottom row
    for j in range(r, l - 1, -1):
        arr.append(matrix[b][j])

    # left column
    for i in range(b - 1, t, -1):
        arr.append(matrix[i][l])

    # Step 2: rotate
    n = len(arr)
    k = turns % n
    if k == 0:
        return

    arr = arr[k:] + arr[:k]  # anti-clockwise = left shift

    # Step 3: put back into matrix
    idx = 0

    # top row
    for j in range(l, r + 1):
        matrix[t][j] = arr[idx]
        idx += 1

    # right column
    for i in range(t + 1, b):
        matrix[i][r] = arr[idx]
        idx += 1

    # bottom row
    for j in range(r, l - 1, -1):
        matrix[b][j] = arr[idx]
        idx += 1

    # left column
    for i in range(b - 1, t, -1):
        matrix[i][l] = arr[idx]
        idx += 1


def printMatrix(matrix, m):
    for i in range(m):
        print(" ".join(map(str, matrix[i])))


def matrixRotation(matrix, turns):
    m, n = len(matrix), len(matrix[0])

    if allSameCheck(matrix, m, n):
        printMatrix(matrix, m)
        return

    t, b, l, r = 0, m - 1, 0, n - 1
    while t < b and l < r:
        rotateLayer(matrix, turns, l, r, t, b)
        t += 1
        b -= 1
        l += 1
        r -= 1

    printMatrix(matrix, m)


if __name__ == "__main__":
    first_multiple_input = input().rstrip().split()
    m = int(first_multiple_input[0])
    n = int(first_multiple_input[1])
    r = int(first_multiple_input[2])

    matrix = []

    for _ in range(m):
        matrix.append(list(map(int, input().rstrip().split())))

    matrixRotation(matrix, r)
