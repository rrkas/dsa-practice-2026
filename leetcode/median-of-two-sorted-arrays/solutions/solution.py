class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        newl = sorted([*nums1, *nums2])
        n = len(newl)
        if n % 2 == 0:
            return (newl[n // 2] + newl[n // 2 - 1]) / 2
        return newl[n // 2]
