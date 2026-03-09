# class Solution:
#     def lengthOfLongestSubstring(self, s: str) -> int:
#         seen = set()
#         l = 0
#         maxlen = 0

#         for r in range(len(s)):
#             while s[r] in seen:
#                 seen.remove(s[l])
#                 l += 1

#             seen.add(s[r])
#             maxlen = max(maxlen, r - l + 1)

#         return maxlen


class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        last = {}
        l = 0
        maxlen = 0

        for r, c in enumerate(s):
            if c in last and last[c] >= l:
                l = last[c] + 1

            last[c] = r
            maxlen = max(maxlen, r - l + 1)

        return maxlen
