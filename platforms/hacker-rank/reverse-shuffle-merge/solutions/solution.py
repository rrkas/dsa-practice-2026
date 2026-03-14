#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'reverseShuffleMerge' function below.
#
# The function is expected to return a STRING.
# The function accepts STRING s as parameter.
#


def reverseShuffleMerge(s):
    from collections import Counter

    freq = Counter(s)
    remain = dict(freq)
    need = {k: v // 2 for k, v in freq.items()}
    used = {k: 0 for k, v in freq.items()}

    stack = []

    for c in s[::-1]:
        remain[c] -= 1

        if used[c] == need[c]:
            continue

        while stack:
            top = stack[-1]

            if top <= c:
                break

            if used[top] - 1 + remain[top] < need[top]:
                break

            stack.pop()
            used[top] -= 1

        stack.append(c)
        used[c] += 1

    return "".join(stack)


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")

    s = input()

    result = reverseShuffleMerge(s)

    fptr.write(result + "\n")

    fptr.close()
