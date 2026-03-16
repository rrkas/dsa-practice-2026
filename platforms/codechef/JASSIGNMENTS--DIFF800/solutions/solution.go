package main

import "fmt"

func main() {
	var t, x int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		fmt.Scan(&x)
		if x+3 <= 10 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
