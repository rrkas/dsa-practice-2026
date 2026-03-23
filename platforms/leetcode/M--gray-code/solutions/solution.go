func grayCode(n int) []int {
	res := make([]int, 0, 1<<n)

	for i := 0; i < (1 << n); i++ {
		res = append(res, i^(i>>1))
	}

	return res
}