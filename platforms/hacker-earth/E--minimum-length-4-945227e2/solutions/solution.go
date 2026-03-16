package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(n int, a []int, b []int) int {
	i, j := 0, n-1
	for i <= j {
		if a[i] == b[i] {
			i++
		} else if a[j] == b[j] {
			j--
		} else {
			break
		}
	}
	return j - i + 1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	scanArr := func() []int {
		scanner.Scan()
		fields := strings.Fields(scanner.Text())
		arr := make([]int, 0)
		for _, field := range fields {
			val, _ := strconv.Atoi(field)
			arr = append(arr, val)
		}
		return arr
	}

	for ; t > 0; t-- {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		a := scanArr()
		b := scanArr()
		fmt.Println(solve(n, a, b))
	}
}
