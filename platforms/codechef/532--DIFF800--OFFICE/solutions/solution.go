package main

import "fmt"

func main() {
	var t, x, y int
	fmt.Scan(&t)

	for ; t > 0; t-- {
		fmt.Scan(&x, &y)
		fmt.Println(x*4 + y)
	}
}
