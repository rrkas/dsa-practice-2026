func myPow(x float64, n int) float64 {
	sign := 1
	if n < 0 {
		sign = -1
		n = -n
	}

	var rec func(float64, int) float64
	rec = func(x float64, n int) float64 {
		if x == 0 {
			return 0
		}

		if n == 0 {
			return 1
		}

		if n == 1 {
			return x
		}

		if n%2 == 0 {
			return rec(x*x, n/2)
		} else {
			return x * rec(x, n-1)
		}
	}

	res := rec(x, n)
	if sign < 0 {
		return 1 / res
	}
	return res
}