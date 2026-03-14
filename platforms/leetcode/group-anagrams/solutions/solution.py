from collections import defaultdict


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        res = defaultdict(list)

        for word in strs:
            count = [0] * 26
            for c in word:
                count[ord(c) - ord("a")] += 1

            res[tuple(count)].append(word)

        return list(res.values())


################################################################################

from collections import defaultdict


class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        groups = defaultdict(list)

        for word in strs:
            key = "".join(sorted(word))
            groups[key].append(word)

        return list(groups.values())
