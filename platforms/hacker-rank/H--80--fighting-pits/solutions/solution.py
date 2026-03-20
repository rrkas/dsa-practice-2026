#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'fightingPits' function below.
#
# The function is expected to return an INTEGER_ARRAY.
# The function accepts following parameters:
#  1. INTEGER k
#  2. 2D_INTEGER_ARRAY fighters
#  3. 2D_INTEGER_ARRAY queries
#


def fightingPits(k, fighters, queries):
    teams = [None, [], []]
    for s, t in fighters:
        teams[t].append(s)

    teams[1].sort()
    teams[2].sort()

    result = []

    for choice, v1, v2 in queries:
        if choice == 1:
            teams[v2].append(v1)
            teams[v2].sort()
        else:
            A = teams[v1]
            B = teams[v2]

            i, j = len(A) - 1, len(B) - 1

            turnA = True
            while i >= 0 and j >= 0:
                if turnA:
                    j -= A[i]
                else:
                    i -= B[j]

                turnA = not turnA

            if j < 0:
                result.append(v1)
            else:
                result.append(v2)

    return result


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")

    first_multiple_input = input().rstrip().split()

    n = int(first_multiple_input[0])

    k = int(first_multiple_input[1])

    q = int(first_multiple_input[2])

    fighters = []

    for _ in range(n):
        fighters.append(list(map(int, input().rstrip().split())))

    queries = []

    for _ in range(q):
        queries.append(list(map(int, input().rstrip().split())))

    result = fightingPits(k, fighters, queries)

    fptr.write("\n".join(map(str, result)))
    fptr.write("\n")

    fptr.close()
