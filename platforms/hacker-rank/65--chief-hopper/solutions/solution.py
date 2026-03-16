#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'chiefHopper' function below.
#
# The function is expected to return an INTEGER.
# The function accepts INTEGER_ARRAY arr as parameter.
#


# def chiefHopper(arr):
#     minEnergy = 1

#     while True:
#         botEnergy = minEnergy
#         complete = True
#         for i, e in enumerate(arr):
#             botEnergy = 2 * botEnergy - e
#             if botEnergy < 0:
#                 complete = False
#                 break

#         if complete:
#             return minEnergy
#         else:
#             minEnergy += 1


def chiefHopper(arr):
    energy = 0

    for h in reversed(arr):
        energy = (energy + h + 1) // 2

    return energy


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    n = int(input().strip())
    arr = list(map(int, input().rstrip().split()))
    result = chiefHopper(arr)
    fptr.write(str(result) + "\n")
    fptr.close()
