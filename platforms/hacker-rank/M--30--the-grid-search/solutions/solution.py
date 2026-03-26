#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'gridSearch' function below.
#
# The function is expected to return a STRING.
# The function accepts following parameters:
#  1. STRING_ARRAY G
#  2. STRING_ARRAY P
#


def gridSearch(G, P):
    M, N = len(G), len(G[0])
    m, n = len(P), len(P[0])

    for i in range(M - m + 1):
        start = 0

        while True:
            j = 0

            try:
                j = G[i][start:].index(P[0])
            except:
                break

            j += start
            found = True

            for k in range(m):
                if G[i + k][j : j + n] != P[k]:
                    found = False
                    break

            if found:
                return "YES"

            start = j + 1

    return "NO"


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    t = int(input().strip())

    for t_itr in range(t):
        first_multiple_input = input().rstrip().split()
        R = int(first_multiple_input[0])
        C = int(first_multiple_input[1])
        G = []

        for _ in range(R):
            G_item = input()
            G.append(G_item)

        second_multiple_input = input().rstrip().split()
        r = int(second_multiple_input[0])
        c = int(second_multiple_input[1])
        P = []

        for _ in range(r):
            P_item = input()
            P.append(P_item)

        result = gridSearch(G, P)
        fptr.write(result + "\n")

    fptr.close()
