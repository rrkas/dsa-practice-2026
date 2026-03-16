t = int(input())

for _ in range(t):
    n, x = map(int, input().split())
    n = (n + 5) // 6
    print(n * x)
