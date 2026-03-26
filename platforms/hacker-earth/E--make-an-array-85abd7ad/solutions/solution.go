package main

import "fmt"

func solve(N int, A []int) int {
	sum := 0
	maxVal := 0

	for i := 0; i < N; i++ {
		sum += A[i]
		if A[i] > maxVal {
			maxVal = A[i]
		}
	}

	if sum%(N-1) != 0 {
		return -1
	}

	k := sum / (N - 1)

	if maxVal > k {
		return -1
	}

	return k
}

func main() {
	var T int
	fmt.Scanln(&T)
	for i := 0; i < T; i++ {

		var N int
		fmt.Scanln(&N)
		A := make([]int, N)
		for i_A := 0; i_A < N; i_A++ {
			fmt.Scan(&A[i_A])
		}

		var out_ int = solve(N, A)
		fmt.Print(out_)
		fmt.Println()
	}
}
