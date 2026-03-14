#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'boardCutting' function below.
#
# The function is expected to return an INTEGER.
# The function accepts following parameters:
#  1. INTEGER_ARRAY cost_y
#  2. INTEGER_ARRAY cost_x
#

MOD = 10**9 + 7


def boardCutting(cost_y, cost_x):
    cost_x.sort(reverse=True)
    cost_y.sort(reverse=True)

    cost = 0
    i, j = 0, 0
    h, v = 1, 1
    while i < len(cost_x) and j < len(cost_y):
        if cost_x[i] > cost_y[j]:
            cost += cost_x[i] * h
            v += 1
            i += 1
        else:
            cost += cost_y[j] * v
            h += 1
            j += 1

        cost %= MOD

    while i < len(cost_x):
        cost += cost_x[i] * h
        v += 1
        i += 1

        cost %= MOD

    while j < len(cost_y):
        cost += cost_y[j] * v
        h += 1
        j += 1

        cost %= MOD

    return cost


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")

    q = int(input().strip())

    for q_itr in range(q):
        first_multiple_input = input().rstrip().split()

        m = int(first_multiple_input[0])

        n = int(first_multiple_input[1])

        cost_y = list(map(int, input().rstrip().split()))

        cost_x = list(map(int, input().rstrip().split()))

        result = boardCutting(cost_y, cost_x)

        fptr.write(str(result) + "\n")

    fptr.close()
