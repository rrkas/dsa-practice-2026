package main

import "fmt"

func main() {
	var t, x, y, z int
	fmt.Scan(&t)
	for ; t > 0; t-- {
		fmt.Scan(&x, &y, &z)
		if z*2 > x*y {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
