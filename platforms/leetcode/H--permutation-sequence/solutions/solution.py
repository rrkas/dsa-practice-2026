class Solution:
    def getPermutation(self, n: int, k: int) -> str:
        def fact(n):
            f = 1
            for i in range(2, n + 1):
                f *= i
            return f

        k -= 1
        nums = list(range(1, n + 1))
        res = ""
        for i in range(n, 0, -1):
            f = fact(i - 1)
            idx = k // f
            nidx = nums.pop(idx)
            res += str(nidx)
            print(f"i={i} f={f} k={k} idx={idx} nums[idx]={nidx} res={res}")
            k %= f

        return res
