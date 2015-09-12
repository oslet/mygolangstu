package main

import "fmt"

func fibonacci() func(int) int {
	sum1 := 0
	sum2 := 1
	return func(x int) int {
		out := sum1 + sum2
		sum1 = sum2
		sum2 = out
		return out
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 40; i++ {
		fmt.Println(f(i))
	}
}
