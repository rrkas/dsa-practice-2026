#!/bin/python3

import math
import os
import random
import re
import sys


# Complete the flatlandSpaceStations function below.
def flatlandSpaceStations(n, c):
    c.sort()

    # distance before first station
    maxd = c[0]

    # distance between stations
    for i in range(1, len(c)):
        maxd = max(maxd, (c[i] - c[i - 1]) // 2)

    # distance after last station
    maxd = max(maxd, (n - 1) - c[-1])

    return maxd


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    nm = input().split()
    n = int(nm[0])
    m = int(nm[1])
    c = list(map(int, input().rstrip().split()))
    result = flatlandSpaceStations(n, c)
    fptr.write(str(result) + "\n")
    fptr.close()
