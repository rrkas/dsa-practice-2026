#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'migratoryBirds' function below.
#
# The function is expected to return an INTEGER.
# The function accepts INTEGER_ARRAY arr as parameter.
#


def migratoryBirds(arr):
    freq = [0] * 5

    maxv = 0
    for e in arr:
        freq[e - 1] += 1
        maxv = max(maxv, freq[e - 1])

    for i, e in enumerate(freq):
        if e == maxv:
            return i + 1


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    arr_count = int(input().strip())
    arr = list(map(int, input().rstrip().split()))
    result = migratoryBirds(arr)
    fptr.write(str(result) + "\n")
    fptr.close()
