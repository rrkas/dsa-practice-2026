class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        hn, nn = len(haystack), len(needle)
        if hn == 0:
            return -1

        if haystack == needle:
            return 0

        for hi in range(hn - nn + 1):
            match = True
            for ni in range(nn):
                if haystack[hi + ni] != needle[ni]:
                    match = False
                    break
            if match:
                return hi
        return -1
