#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'plusMinus' function below.
#
# The function accepts INTEGER_ARRAY arr as parameter.
#


def plusMinus(arr):
    n = len(arr)
    pos = sum([e > 0 for e in arr])
    neg = sum([e < 0 for e in arr])
    zero = n - pos - neg
    print(f"{pos/n:.6f}")
    print(f"{neg/n:.6f}")
    print(f"{zero/n:.6f}")


if __name__ == "__main__":
    n = int(input().strip())

    arr = list(map(int, input().rstrip().split()))

    plusMinus(arr)
