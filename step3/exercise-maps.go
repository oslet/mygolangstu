package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)

	arr := strings.Fields(s)
	for _, val := range arr {
		ret[val]++
	}
	return ret

}

func main() {
	wc.Test(WordCount)
}
