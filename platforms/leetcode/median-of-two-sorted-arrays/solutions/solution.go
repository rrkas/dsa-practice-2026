import "slices"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    comb := append(nums1, nums2...)
    slices.Sort(comb)

    n := len(comb)

    if n%2 == 0 {
        return float64(comb[n/2] + comb[n/2-1]) / 2.0
    }

    return float64(comb[n/2])
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    m, n := len(nums1), len(nums2)
    total := m + n

    i, j := 0, 0
    prev, curr := 0, 0

    for k := 0; k <= total/2; k++ {
        prev = curr

        if i < m && (j >= n || nums1[i] <= nums2[j]) {
            curr = nums1[i]
            i++
        } else {
            curr = nums2[j]
            j++
        }
    }

    if total%2 == 0 {
        return float64(prev+curr) / 2.0
    }
    return float64(curr)
}