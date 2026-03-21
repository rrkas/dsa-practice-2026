#!/bin/python3

import math
import os
import random
import re
import sys
from collections import Counter

#
# Complete the 'pickingNumbers' function below.
#
# The function is expected to return an INTEGER.
# The function accepts INTEGER_ARRAY a as parameter.
#


def pickingNumbers(a):
    freq = Counter()
    for e in a:
        freq[e] += 1

    maxv = 0
    for k, v in freq.items():
        if k + 1 in freq:
            maxv = max(maxv, v + freq[k + 1])
        else:
            maxv = max(maxv, v)

    return maxv


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    n = int(input().strip())
    a = list(map(int, input().rstrip().split()))
    result = pickingNumbers(a)
    fptr.write(str(result) + "\n")
    fptr.close()
