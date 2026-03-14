#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'pylons' function below.
#
# The function is expected to return an INTEGER.
# The function accepts following parameters:
#  1. INTEGER k
#  2. INTEGER_ARRAY arr
#


def pylons(k, arr):
    n = len(arr)
    i = 0
    plants = 0

    while i < n:
        pos = min(n - 1, i + k - 1)

        while pos >= i - (k - 1) and pos >= 0 and arr[pos] == 0:
            pos -= 1

        if pos < i - (k - 1) or pos < 0:
            return -1

        plants += 1

        i = pos + k

    return plants


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")

    first_multiple_input = input().rstrip().split()

    n = int(first_multiple_input[0])

    k = int(first_multiple_input[1])

    arr = list(map(int, input().rstrip().split()))

    result = pylons(k, arr)

    fptr.write(str(result) + "\n")

    fptr.close()
