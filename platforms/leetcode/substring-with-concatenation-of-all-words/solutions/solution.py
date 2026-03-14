from collections import Counter


class Solution:
    def findSubstring(self, s: str, words: List[str]) -> List[int]:
        if len(s) == 0 or len(words) == 0:
            return []

        res = []
        wc = Counter()
        for w in words:
            wc[w] += 1

        wcount = len(words)
        wlen = len(words[0])
        slen = len(s)
        total = wlen * wcount

        for i in range(slen - total + 1):
            seen = Counter()

            j = 0
            while j < wcount:
                start = i + (j * wlen)
                ss = s[start : start + wlen]
                if ss not in wc:
                    break
                seen[ss] += 1
                if seen[ss] > wc[ss]:
                    break
                j += 1

            if j == wcount:
                res.append(i)

        return res
