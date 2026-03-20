#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'dayOfProgrammer' function below.
#
# The function is expected to return a STRING.
# The function accepts INTEGER year as parameter.
#


def getDays(year):
    days = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
    if year == 1918:
        days[1] -= 13
    elif year < 1918:
        if year % 4 == 0:
            days[1] = 29
    else:
        if (year % 400 == 0) or ((year % 4 == 0) and (year % 100 != 0)):
            days[1] = 29
    return days


def dayOfProgrammer(year):
    days = getDays(year)
    if sum(days) == 352:
        return "26.09.1918"
    elif sum(days) == 366:
        return f"12.09.{year}"
    else:
        return f"13.09.{year}"


if __name__ == "__main__":
    fptr = open(os.environ["OUTPUT_PATH"], "w")
    year = int(input().strip())
    result = dayOfProgrammer(year)
    fptr.write(result + "\n")
    fptr.close()
