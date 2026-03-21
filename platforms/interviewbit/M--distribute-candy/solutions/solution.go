/**
 * @input A : Integer array
 *
 * @Output Integer
 */
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func candy(A []int) int {
	n := len(A)
	if n == 0 {
		return 0
	}

	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
	}

	for i := 1; i < n; i++ {
		if A[i] > A[i-1] {
			res[i] = res[i-1] + 1
		}
	}

	sum := res[n-1]
	for i := n - 2; i >= 0; i-- {
		if A[i] > A[i+1] {
			res[i] = max(res[i], res[i+1]+1)
		}
		sum += res[i]
	}
	return sum
}
