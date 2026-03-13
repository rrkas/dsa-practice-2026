class Solution:
    def insert(self, intervals: List[List[int]], new: List[int]) -> List[List[int]]:
        res = []
        i = 0
        n = len(intervals)

        # 1. add before intervals
        while i < n and intervals[i][1] < new[0]:
            res.append(intervals[i])
            i += 1

        # 2. merge overlaps
        while i < n and intervals[i][0] <= new[1]:
            new[0] = min(new[0], intervals[i][0])
            new[1] = max(new[1], intervals[i][1])
            i += 1

        res.append(new)

        # 3. add remaining
        while i < n:
            res.append(intervals[i])
            i += 1

        return res
