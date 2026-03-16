import "slices"

func permuteUnique(nums []int) [][]int {
    slices.Sort(nums)
    
    n := len(nums)
    res := make([][]int, 0)
    seen := make([]bool, n)

    var perm func([]int)
    perm = func(path []int) {
        if len(path) == n {
            res = append(res, append([]int{}, path...))
            return
        }

        for i := 0; i < n; i++ {
            if seen[i] {
                continue
            }

            if i > 0 && nums[i] == nums[i-1] && seen[i-1] {
                continue
            }

            seen[i] = true
            path = append(path, nums[i])
            perm(path)
            path = path[:len(path)-1]
            seen[i] = false
        }
    }

    perm([]int{})
    return res    
}

