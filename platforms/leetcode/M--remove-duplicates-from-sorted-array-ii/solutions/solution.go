func removeDuplicates(nums []int) int {
    n := len(nums)
    if n <= 2 {
        return n
    }

    k := 2 // first 2 always valid

    for i := 2; i < n; i++ {
        // fmt.Println(nums, i, k)
        if nums[i] != nums[k-2] {
            nums[k] = nums[i]
            k++
        }
    }

    return k
}