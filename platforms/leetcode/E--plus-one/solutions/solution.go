func plusOne(digits []int) []int {
	carry := 0
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		v := digits[i] + carry
		if i == n-1 {
			v++
		}
		digits[i] = v % 10
		carry = v / 10
	}
	if carry > 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}