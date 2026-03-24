#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'happyLadybugs' function below.
#
# The function is expected to return a STRING.
# The function accepts STRING b as parameter.
#


def happyLadybugs(b):
    # frequency check
    from collections import Counter

    freq = Counter()
    for c in b:
        if c != "_":
            freq[c] += 1
    for v in freq.values():
        if v == 1:
            return "NO"

    # adjacency & vancancy
    vacancies = b.count("_")
    if vacancies > 0:
        return "YES"

    for i in range(len(b)):
        f = False
        if i > 0 and b[i] == b[i - 1]:
            f = True
        if i + 1 < len(b) and b[i] == b[i + 1]:
            f = True
        if not f:
            return "NO"

    return "YES"


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    g = int(input().strip())

    for g_itr in range(g):
        n = int(input().strip())
        b = input()
        result = happyLadybugs(b)
        fptr.write(result + "\n")

    fptr.close()
