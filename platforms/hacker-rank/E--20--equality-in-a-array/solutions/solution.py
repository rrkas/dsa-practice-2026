#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'equalizeArray' function below.
#
# The function is expected to return an INTEGER.
# The function accepts INTEGER_ARRAY arr as parameter.
#


def equalizeArray(arr):
    from collections import Counter

    freq = Counter(arr)
    maxv = max(freq.values())
    return len(arr) - maxv


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    n = int(input().strip())
    arr = list(map(int, input().rstrip().split()))
    result = equalizeArray(arr)
    fptr.write(str(result) + "\n")
    fptr.close()
