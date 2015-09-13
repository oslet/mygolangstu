package main

import (
	"fmt"
	"math"
)

func prime(value int) bool {
	if value <= 1 {
		return false
	}
	if value == 2 || value == 3 || value == 5 || value == 7 {
		return true
	}
	if value%2 == 0 || value%3 == 0 || value%5 == 0 || value%7 == 0 {
		return false
	}
	factor := 7
	c := []int{4, 2, 4, 2, 4, 6, 2, 6}
	max := int(math.Sqrt(float64(value)))
	if max*max == value {
		return false
	}
	for factor < max {
		for i := 0; i < len(c); i++ {
			factor += c[i]
			if value%factor == 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	var nCount int
	n := 1000
	for i := 1; i <= n; i++ {
		if prime(i) {
			nCount += 1
			fmt.Printf("%5d", i)
			if nCount%8 == 0 {
				fmt.Println("")
			}
		}
	}
	fmt.Println("Count =", nCount)
}
