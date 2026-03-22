#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'viralAdvertising' function below.
#
# The function is expected to return an INTEGER.
# The function accepts INTEGER n as parameter.
#


def viralAdvertising(n):
    shared = 5
    likes = 0
    cum = 0
    for i in range(n):
        likes = shared // 2
        cum += likes
        shared = likes * 3
    return cum


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    n = int(input().strip())
    result = viralAdvertising(n)
    fptr.write(str(result) + "\n")
    fptr.close()
