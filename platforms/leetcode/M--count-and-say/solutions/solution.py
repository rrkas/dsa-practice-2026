class Solution:
    def countAndSay(self, n: int) -> str:
        def rle(ins: str) -> str:
            t = ""
            i = 0
            while i < len(ins):
                char = s[i]
                count = 1
                while i + 1 < len(ins) and s[i] == s[i + 1]:
                    i += 1
                    count += 1
                t += str(count) + char
                i += 1
            return t

        if n == 0:
            return ""

        s = "1"
        for _ in range(1, n):
            s = rle(s)

        return s
