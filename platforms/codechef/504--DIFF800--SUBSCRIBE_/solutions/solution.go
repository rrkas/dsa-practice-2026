package main

import "fmt"

func main() {
	var t, n, x int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		fmt.Scan(&n, &x)
		n = (n + 5) / 6
		fmt.Println(n * x)
	}
}
