package main

import "fmt"

func fib(n uint) (ret int64) {
	var index uint = 0
	var i int64 = 1
	var j int64 = 0

	for index < n {
		index += 1
		i, j = i+j, i
	}
	return j
}

func main() {
	fmt.Printf("%d\n", fib(1000))
}
