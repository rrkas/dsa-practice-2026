#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'timeConversion' function below.
#
# The function is expected to return a STRING.
# The function accepts STRING s as parameter.
#

def timeConversion(s):
    is_am = s[8] == "A"
    hh = int(s[:2])
    if not is_am and hh != 12:
        hh += 12
    if is_am and hh == 12:
        hh = 0
    
    return f"{hh:02d}" + s[2:8]

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = input()

    result = timeConversion(s)

    fptr.write(result + '\n')

    fptr.close()
