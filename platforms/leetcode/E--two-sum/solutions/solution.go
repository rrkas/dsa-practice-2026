// func twoSum(nums []int, target int) []int {
//     n := len(nums)
//     for i := 0; i<n-1; i++ {
//         for j := i+1; j < n; j++ {
//             if nums[i] + nums[j] == target {
//                 return []int{i, j}
//             }
//         }
//     }
//     return []int{}
// }

func twoSum(nums []int, target int) []int {
    m := make(map[int]int)

    for i, v := range nums {
        complement := target - v

        if idx, ok := m[complement]; ok {
            return []int{idx, i}
        }

        m[v] = i
    }

    return []int{}
}
