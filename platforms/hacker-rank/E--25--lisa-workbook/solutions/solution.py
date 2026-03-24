#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'workbook' function below.
#
# The function is expected to return an INTEGER.
# The function accepts following parameters:
#  1. INTEGER n
#  2. INTEGER k
#  3. INTEGER_ARRAY arr
#


def workbook(n, k, arr):
    count = 0

    page = 1
    for e in arr:
        pages = (e + k - 1) // k
        for p in range(pages):
            qs, qe = p * k + 1, min((p + 1) * k, e)
            if qs <= page <= qe:
                count += 1
            page += 1

    return count


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    first_multiple_input = input().rstrip().split()
    n = int(first_multiple_input[0])
    k = int(first_multiple_input[1])
    arr = list(map(int, input().rstrip().split()))
    result = workbook(n, k, arr)
    fptr.write(str(result) + "\n")
    fptr.close()
