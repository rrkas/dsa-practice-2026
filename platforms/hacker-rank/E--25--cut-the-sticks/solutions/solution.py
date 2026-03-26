#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'cutTheSticks' function below.
#
# The function is expected to return an INTEGER_ARRAY.
# The function accepts INTEGER_ARRAY arr as parameter.
#


def cutTheSticks(arr):
    arr.sort()
    n = len(arr)
    res = []
    for i in range(n):
        if i == 0 or arr[i] != arr[i - 1]:
            res.append(n - i)

    return res


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    n = int(input().strip())
    arr = list(map(int, input().rstrip().split()))
    result = cutTheSticks(arr)
    fptr.write("\n".join(map(str, result)))
    fptr.write("\n")
    fptr.close()
