import "strings"

func getPermutation(n int, k int) string {
	res := strings.Builder{}

	k--
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i + 1
	}

	for i := n; i > 0; i-- {
		f := fact(i - 1)
		idx := k / f
		num := nums[idx]
		nums = append(nums[:idx], nums[idx+1:]...)
		res.WriteByte('0' + byte(num))
		k %= f
	}

	return res.String()
}

func fact(n int) int {
	f := 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}
