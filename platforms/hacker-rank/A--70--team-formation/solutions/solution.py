import heapq


def solve_case(skills):
    if not skills:
        return 0

    skills.sort()
    end = {}
    ans = float("inf")

    for x in skills:
        if x - 1 in end and end[x - 1]:
            length = heapq.heappop(end[x - 1])
            if not end[x - 1]:
                del end[x - 1]

            end.setdefault(x, [])
            heapq.heappush(end[x], length + 1)
        else:
            end.setdefault(x, [])
            heapq.heappush(end[x], 1)

    for heap in end.values():
        ans = min(ans, min(heap))

    return ans


t = int(input())
for _ in range(t):
    arr = list(map(int, input().split()))
    n = arr[0]
    skills = arr[1:]
    print(solve_case(skills))
