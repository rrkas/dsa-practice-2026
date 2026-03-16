def solve(n, a, b):
    i, j = 0, n - 1
    while i <= j:
        if a[i] == b[i]:
            i += 1
        elif a[j] == b[j]:
            j -= 1
        else:
            break
    return j - i + 1


t = int(input())
for _ in range(t):
    n = int(input())
    a = list(map(int, input().split()))
    b = list(map(int, input().split()))
    print(solve(n, a, b))
