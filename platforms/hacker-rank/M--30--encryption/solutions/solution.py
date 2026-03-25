#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'encryption' function below.
#
# The function is expected to return a STRING.
# The function accepts STRING s as parameter.
#


def encryption(s):
    sq = len(s) ** 0.5
    m, n = math.floor(sq), math.ceil(sq)
    if m * n < len(s):
        m += 1

    res = ""
    for j in range(n):
        for i in range(m):
            si = i * n + j
            if si >= len(s):
                continue
            res += s[si]
        res += " "

    return res.strip()


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    s = input()
    result = encryption(s)
    fptr.write(result + "\n")
    fptr.close()
