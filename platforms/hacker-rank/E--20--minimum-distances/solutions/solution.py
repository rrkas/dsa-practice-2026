#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'minimumDistances' function below.
#
# The function is expected to return an INTEGER.
# The function accepts INTEGER_ARRAY a as parameter.
#


def minimumDistances(a):
    mind = len(a)

    seen = {}

    for i, e in enumerate(a):
        if e not in seen:
            seen[e] = i
        else:
            mind = min(mind, i - seen[e])

    if mind == len(a):
        return -1

    return mind


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    n = int(input().strip())
    a = list(map(int, input().rstrip().split()))
    result = minimumDistances(a)
    fptr.write(str(result) + "\n")
    fptr.close()
