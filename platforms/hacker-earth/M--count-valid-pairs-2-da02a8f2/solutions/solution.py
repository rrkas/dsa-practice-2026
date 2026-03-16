from collections import Counter


def numsum(n):
    s = 0
    while n > 0:
        s += n % 10
        n //= 10
    return s


def solve(n, nums):
    count = Counter()

    for e in nums:
        count[numsum(e)] += 1

    res = 0
    for v in count.values():
        res += v * (v - 1) // 2
    return res


n = int(input())
nums = list(map(int, input().split()))

out_ = solve(n, nums)
print(out_)
