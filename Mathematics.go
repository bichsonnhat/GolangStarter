package main

import (
	"fmt"
)

const MOD = 1e9 + 7

func main() {
	var n int
	fmt.Scan(&n)
	var opr bool = false
	var a []int
	for i := 0; i < n; i++ {
		var u int
		fmt.Scan(&u)
		a = append(a, u)
	}
	for condition := true; condition; condition = (len(a) != 1) {
		if opr == false {
			for i := 0; i+1 < len(a); i++ {
				a[i] += a[i+1]
				a[i] %= MOD
			}
			a = a[:len(a)-1]
			opr = true
			continue
		}
		if opr == true {
			for i := 0; i+1 < len(a); i++ {
				a[i] *= a[i+1]
				a[i] %= MOD
			}
			a = a[:len(a)-1]
			opr = false
			continue
		}
	}
	fmt.Print(a[0])
}
