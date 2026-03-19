#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'breakingRecords' function below.
#
# The function is expected to return an INTEGER_ARRAY.
# The function accepts INTEGER_ARRAY scores as parameter.
#


def breakingRecords(scores):
    min_, max_ = scores[0], scores[0]
    minc, maxc = 0, 0

    for e in scores[1:]:
        if e < min_:
            min_ = e
            minc += 1
        if e > max_:
            max_ = e
            maxc += 1
    return maxc, minc


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    n = int(input().strip())
    scores = list(map(int, input().rstrip().split()))
    result = breakingRecords(scores)
    fptr.write(" ".join(map(str, result)))
    fptr.write("\n")
    fptr.close()
