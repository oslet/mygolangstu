package main

import "fmt"

func fib(n uint) uint64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	fmt.Printf("%d\n", fib(40))
}
