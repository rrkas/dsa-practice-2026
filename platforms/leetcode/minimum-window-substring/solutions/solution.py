from collections import Counter


class Solution:
    def minWindow(self, s: str, t: str) -> str:
        if len(t) > len(s):
            return ""

        if t == s:
            return s

        tmap = Counter(t)

        l = 0
        minlen = len(s) + 1
        start = 0

        window = Counter()
        required = len(tmap)
        formed = 0

        for r in range(len(s)):
            c = s[r]
            window[c] += 1

            if window[c] == tmap[c]:
                formed += 1

            while l <= r and formed == required:
                if r - l + 1 < minlen:
                    minlen = r - l + 1
                    start = l

                window[s[l]] -= 1
                if window[s[l]] < tmap[s[l]]:
                    formed -= 1

                l += 1

        if minlen == len(s) + 1:
            return ""

        return s[start : start + minlen]
