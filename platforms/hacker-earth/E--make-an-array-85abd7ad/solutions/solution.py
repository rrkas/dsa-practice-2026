def solve(N, A):
    s = sum(A)
    m = max(A)
    if s % (N - 1) != 0:
        return -1

    k = s // (N - 1)
    if m > k:
        return -1

    return k


T = int(input())
for _ in range(T):
    N = int(input())
    A = list(map(int, input().split()))

    out_ = solve(N, A)
    print(out_)
