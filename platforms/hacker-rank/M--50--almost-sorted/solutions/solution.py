#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'almostSorted' function below.
#
# The function accepts INTEGER_ARRAY arr as parameter.
#


def almostSorted(arr):
    n = len(arr)

    # Step 1: find first and last violation
    l, r = -1, -1

    for i in range(n - 1):
        if arr[i] > arr[i + 1]:
            if l == -1:
                l = i
            r = i + 1

    # already sorted
    if l == -1:
        print("yes")
        return

    # Step 2: try swap
    arr[l], arr[r] = arr[r], arr[l]

    if all(arr[i] <= arr[i + 1] for i in range(n - 1)):
        print("yes")
        print(f"swap {l+1} {r+1}")
        return

    # undo swap
    arr[l], arr[r] = arr[r], arr[l]

    # Step 3: try reverse
    arr[l : r + 1] = arr[l : r + 1][::-1]

    if all(arr[i] <= arr[i + 1] for i in range(n - 1)):
        print("yes")
        print(f"reverse {l+1} {r+1}")
        return

    # Step 4: not possible
    print("no")


if __name__ == "__main__":
    n = int(input().strip())

    arr = list(map(int, input().rstrip().split()))

    almostSorted(arr)
