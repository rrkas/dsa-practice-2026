#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'kaprekarNumbers' function below.
#
# The function accepts following parameters:
#  1. INTEGER p
#  2. INTEGER q
#


def checkKaprekar(n):
    numdigs = len(str(n))
    if n == 1:
        return True

    sq = str(n**2)

    if numdigs == 1 and len(sq) == 1:
        return False

    l, r = int(sq[:-numdigs]), int(sq[-numdigs:])
    return l + r == n


def kaprekarNumbers(p, q):
    res = []
    for i in range(p, q + 1):
        if checkKaprekar(i):
            res.append(i)
    if res:
        print(" ".join(map(str, res)))
    else:
        print("INVALID RANGE")


if __name__ == "__main__":
    p = int(input().strip())
    q = int(input().strip())
    kaprekarNumbers(p, q)
