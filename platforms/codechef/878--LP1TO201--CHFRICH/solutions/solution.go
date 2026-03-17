package main

import "fmt"

func main() {
	var t, a, b, x int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		fmt.Scan(&a, &b, &x)
		fmt.Println((b - a) / x)
	}
}
