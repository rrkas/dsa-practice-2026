#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'appendAndDelete' function below.
#
# The function is expected to return a STRING.
# The function accepts following parameters:
#  1. STRING s
#  2. STRING t
#  3. INTEGER k
#


def appendAndDelete(s, t, k):
    sn, tn = len(s), len(t)

    if k >= sn + tn:
        return "Yes"

    plen = 0

    for i in range(min(sn, tn)):
        if s[i] == t[i]:
            plen += 1
        else:
            break

    minops = sn + tn - 2 * plen
    if minops > k:
        return "No"

    if minops == k:
        return "Yes"

    if (minops - k) % 2 == 0:
        return "Yes"

    return "No"


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    s = input()
    t = input()
    k = int(input().strip())
    result = appendAndDelete(s, t, k)
    fptr.write(result + "\n")
    fptr.close()
