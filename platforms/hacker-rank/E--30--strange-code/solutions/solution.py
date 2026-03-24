#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'strangeCounter' function below.
#
# The function is expected to return a LONG_INTEGER.
# The function accepts LONG_INTEGER t as parameter.
#


def strangeCounter(t):
    start = 3

    while start < t:
        t -= start
        start *= 2

    return start - t + 1


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    t = int(input().strip())
    result = strangeCounter(t)
    fptr.write(str(result) + "\n")
    fptr.close()
