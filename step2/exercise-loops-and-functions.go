package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	if z := math.Sqrt(8); x < z {
		return z
	}
	return x
}

func main() {
	fmt.Println(sqrt(2))
}
