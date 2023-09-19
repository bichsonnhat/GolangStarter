package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var a, b int64
		fmt.Scan(&a, &b)
		fmt.Print(a+b, "\n")
	}
}
