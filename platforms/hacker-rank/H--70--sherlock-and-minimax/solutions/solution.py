#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'sherlockAndMinimax' function below.
#
# The function is expected to return an INTEGER.
# The function accepts following parameters:
#  1. INTEGER_ARRAY arr
#  2. INTEGER p
#  3. INTEGER q
#


def minDist(x, arr):
    m = abs(x - arr[0])
    for e in arr[1:]:
        m = min(m, abs(x - e))
    return m


def sherlockAndMinimax(arr, p, q):
    arr.sort()

    bestDist, bestVal = -1, p

    def check(v):
        nonlocal bestDist, bestVal
        if v < p or v > q:
            return
        d = minDist(v, arr)
        if d > bestDist or (d == bestDist and v < bestVal):
            bestDist = d
            bestVal = v

    check(p)
    check(q)

    for i in range(len(arr) - 1):
        mid = (arr[i] + arr[i + 1]) // 2
        check(mid)
        check(mid + 1)

    return bestVal


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    n = int(input().strip())
    arr = list(map(int, input().rstrip().split()))
    first_multiple_input = input().rstrip().split()
    p = int(first_multiple_input[0])
    q = int(first_multiple_input[1])
    result = sherlockAndMinimax(arr, p, q)
    fptr.write(str(result) + "\n")
    fptr.close()
