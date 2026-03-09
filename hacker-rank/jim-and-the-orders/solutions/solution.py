#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'jimOrders' function below.
#
# The function is expected to return an INTEGER_ARRAY.
# The function accepts 2D_INTEGER_ARRAY orders as parameter.
#


def jimOrders(orders):
    class Entry:
        def __init__(self, cid, serveTime):
            self.cid = cid
            self.serveTime = serveTime

    entries = []
    for cid, (orderTime, prepTime) in enumerate(orders):
        entries.append(Entry(cid + 1, orderTime + prepTime))

    entries.sort(key=lambda e: (e.serveTime, e.cid))
    return [e.cid for e in entries]

    # Write your code here


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")

    n = int(input().strip())

    orders = []

    for _ in range(n):
        orders.append(list(map(int, input().rstrip().split())))

    result = jimOrders(orders)

    fptr.write(" ".join(map(str, result)))
    fptr.write("\n")

    fptr.close()
