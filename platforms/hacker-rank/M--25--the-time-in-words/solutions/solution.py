#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'timeInWords' function below.
#
# The function is expected to return a STRING.
# The function accepts following parameters:
#  1. INTEGER h
#  2. INTEGER m
#


def numToWords(n):
    if n >= 20:
        tens = n // 10
        tensW = {
            2: "twenty",
            3: "thirty",
            4: "forty",
            5: "fifty",
        }[tens]
        return tensW + " " + numToWords(n % 10)

    return {
        1: "one",
        2: "two",
        3: "three",
        4: "four",
        5: "five",
        6: "six",
        7: "seven",
        8: "eight",
        9: "nine",
        10: "ten",
        11: "eleven",
        12: "twelve",
        13: "thirteen",
        14: "fourteen",
        15: "fifteen",
        16: "sixteen",
        17: "seventeen",
        18: "eighteen",
        19: "nineteen",
    }[n]


def timeInWords(h, m):
    match m:
        case 0:
            return numToWords(h) + " o' clock"
        case 15:
            return "quarter past " + numToWords(h)
        case 30:
            return "half past " + numToWords(h)
        case 45:
            return "quarter to " + numToWords(h + 1)
        case 1:
            return "one minute past " + numToWords(h)
        case 59:
            return "one minute to " + numToWords(h + 1)
        case _:
            if m < 30:
                return numToWords(m) + " minutes past " + numToWords(h)
            else:
                return numToWords(60 - m) + " minutes to " + numToWords(h + 1)


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    h = int(input().strip())
    m = int(input().strip())
    result = timeInWords(h, m)
    fptr.write(result + "\n")
    fptr.close()
