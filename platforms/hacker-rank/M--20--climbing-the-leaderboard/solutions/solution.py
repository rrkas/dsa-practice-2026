#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'climbingLeaderboard' function below.
#
# The function is expected to return an INTEGER_ARRAY.
# The function accepts following parameters:
#  1. INTEGER_ARRAY ranked
#  2. INTEGER_ARRAY player
#


def climbingLeaderboard(ranked, player):
    # remove duplicates
    unq = []
    for score in ranked:
        if not unq or unq[-1] != score:
            unq.append(score)

    res = []
    i = len(unq) - 1  # start from lowest rank

    for p in player:
        while i >= 0 and p >= unq[i]:
            i -= 1
        res.append(i + 2)

    return res


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    ranked_count = int(input().strip())
    ranked = list(map(int, input().rstrip().split()))
    player_count = int(input().strip())
    player = list(map(int, input().rstrip().split()))
    result = climbingLeaderboard(ranked, player)
    fptr.write("\n".join(map(str, result)))
    fptr.write("\n")
    fptr.close()
