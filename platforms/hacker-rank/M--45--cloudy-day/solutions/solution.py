#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'maximumPeople' function below.
#
# The function is expected to return a LONG_INTEGER.
# The function accepts following parameters:
#  1. LONG_INTEGER_ARRAY p
#  2. LONG_INTEGER_ARRAY x
#  3. LONG_INTEGER_ARRAY y
#  4. LONG_INTEGER_ARRAY r
#


def maximumPeople(p, x, y, r):
    # Return the maximum number of people that will be in a sunny town after removing exactly one cloud.
    n = len(p)
    m = len(y)

    class Event:
        def __init__(self, pos, typ, id, pop):
            self.pos = pos
            self.typ = typ
            self.id = id
            self.pop = pop

    events = []

    for i in range(m):
        events.append(Event(y[i] - r[i], 0, i, 0))
        events.append(Event(y[i] + r[i], 2, i, 0))

    for i in range(n):
        events.append(Event(x[i], 1, i, p[i]))

    events.sort(key=lambda e: (e.pos, e.typ))

    active = set()
    unique = [0] * m
    sunny = 0

    for e in events:
        if e.typ == 0:
            active.add(e.id)
        elif e.typ == 2:
            active.remove(e.id)
        else:
            if len(active) == 0:
                sunny += e.pop
            elif len(active) == 1:
                for id in active:
                    unique[id] += e.pop

    return sunny + max(unique)


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    n = int(input().strip())
    p = list(map(int, input().rstrip().split()))
    x = list(map(int, input().rstrip().split()))
    m = int(input().strip())
    y = list(map(int, input().rstrip().split()))
    r = list(map(int, input().rstrip().split()))
    result = maximumPeople(p, x, y, r)
    fptr.write(str(result) + "\n")
    fptr.close()
