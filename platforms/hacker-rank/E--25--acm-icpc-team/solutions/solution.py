#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'acmTeam' function below.
#
# The function is expected to return an INTEGER_ARRAY.
# The function accepts STRING_ARRAY topic as parameter.
#


def countCommon(s1, s2):
    c = 0
    for i in range(len(s1)):
        c += int(s1[i] == "1" or s2[i] == "1")
    return c


def acmTeam(topic):
    maxc, cmax = 0, 0

    n = len(topic)
    for i in range(n - 1):
        for j in range(i + 1, n):
            c = countCommon(topic[i], topic[j])
            # print(i, j, c)
            if c > maxc:
                maxc = c
                cmax = 1
            elif c == maxc:
                cmax += 1

    return [maxc, cmax]


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    first_multiple_input = input().rstrip().split()
    n = int(first_multiple_input[0])
    m = int(first_multiple_input[1])
    topic = []

    for _ in range(n):
        topic_item = input()
        topic.append(topic_item)

    result = acmTeam(topic)
    fptr.write("\n".join(map(str, result)))
    fptr.write("\n")
    fptr.close()
