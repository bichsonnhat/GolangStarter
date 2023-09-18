package main

import (
	"fmt"
)

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		if max(max(a+b, b+c), a+c) > 10 {
			fmt.Print("YES\n")
		} else {
			fmt.Print("NO\n")
		}
	}
}
