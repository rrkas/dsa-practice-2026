package main

import "fmt"

func main() {
	var t, a, b, c int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		fmt.Scan(&a, &b, &c)
		if a+b > 2*c {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
