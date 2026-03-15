package main

import "fmt"

func numsum(n int) int {
	s := 0
	for n > 0 {
		s += n % 10
		n = n / 10
	}
	return s
}

func solve(n int, nums []int) int64 {
	count := make(map[int]int)
	for _, n := range nums {
		s := numsum(n)
		count[s]++
	}

	pairs := 0
	for _, v := range count {
		pairs += v * (v - 1) / 2
		// fmt.Println(k, v, pairs)
	}
	return int64(pairs)
}

func main() {

	var n int
	fmt.Scanln(&n)
	nums := make([]int, n)
	for i_nums := 0; i_nums < n; i_nums++ {
		fmt.Scan(&nums[i_nums])
	}

	var out_ int64 = solve(n, nums)
	fmt.Print(out_)
}
