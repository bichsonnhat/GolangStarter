package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		var cur int = 0
		var ans int
		for j := 0; j < n; j++ {
			var a, b int
			fmt.Scan(&a, &b)
			if a <= 10 {
				if b > cur {
					ans = j + 1
					cur = b
				}
			}
		}
		fmt.Print(ans, "\n")
	}
}
