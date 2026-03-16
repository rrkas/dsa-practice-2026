def check(a, b, c):
    return "YES" if (a + b) / 2 > c else "NO"


t = int(input())
for _ in range(t):
    a, b, c = map(int, input().split())
    print(check(a, b, c))
