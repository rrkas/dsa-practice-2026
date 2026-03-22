#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'beautifulTriplets' function below.
#
# The function is expected to return an INTEGER.
# The function accepts following parameters:
#  1. INTEGER d
#  2. INTEGER_ARRAY arr
#


def beautifulTriplets(d, arr):
    from collections import Counter

    c = 0

    freq = Counter()
    for e in arr:
        freq[e] += 1

    for k in freq:
        if (k + d) in freq and (k + 2 * d) in freq:
            c += freq[k]

    return c


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    first_multiple_input = input().rstrip().split()
    n = int(first_multiple_input[0])
    d = int(first_multiple_input[1])
    arr = list(map(int, input().rstrip().split()))
    result = beautifulTriplets(d, arr)
    fptr.write(str(result) + "\n")
    fptr.close()
