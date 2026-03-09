class Solution:
    def myAtoi(self, s: str) -> int:
        v = 0
        s = s.strip()

        if len(s) == 0:
            return 0

        sign = -1 if s[0] == "-" else 1

        if s[0] in ["-", "+"]:
            s = s[1:]

        for c in s:
            if c.isdigit():
                v = v * 10 + int(c)
            else:
                break

        v = sign * v

        if v < -(2**31):
            v = -(2**31)

        elif v > 2**31 - 1:
            v = 2**31 - 1

        return v
