func divide(dividend int, divisor int) int {
	const INT_MAX = 1<<31 - 1
	const INT_MIN = -1 << 31

	if dividend == INT_MIN && divisor == -1 {
		return INT_MAX
	}

	sign := 1
	if (dividend < 0) != (divisor < 0) {
		sign = -1
	}

	a := abs(dividend)
	b := abs(divisor)

	result := 0

	for a >= b {
		multiple := 1

		for a >= (b * (multiple << 1)) {
			multiple <<= 1
		}

		a -= b * multiple
		result += multiple
	}

	if sign < 0 {
		return -result
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}