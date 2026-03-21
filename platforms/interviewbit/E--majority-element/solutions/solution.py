class Solution:
    # @param A : tuple of integers
    # @return an integer
    def majorityElement(self, A):
        from collections import Counter

        n = len(A)
        freq = Counter()
        for e in A:
            freq[e] += 1
            if freq[e] > n // 2:
                return e
