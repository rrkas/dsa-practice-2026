#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'staircase' function below.
#
# The function accepts INTEGER n as parameter.
#


def staircase(n):
    for i in range(n):
        for s in range(n - 1 - i):
            print(" ", end="")
        for c in range(i + 1):
            print("#", end="")
        print()


if __name__ == "__main__":
    n = int(input().strip())

    staircase(n)
