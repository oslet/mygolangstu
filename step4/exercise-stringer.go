/*
练习：Stringers
让 IPAddr 类型实现 fmt.Stringer 以便用点分格式输出地址。

例如，IPAddr{1, 2, 3, 4} 应当输出 "1.2.3.4"。
*/

package main

import "fmt"

type IPAddr struct {
	a, b, c, d int
}

func (p IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", p.a, p.b, p.c, p.d)
}

func main() {
	addrs := map[string]IPAddr{
		"loopback":   {127, 0, 0, 1},
		"googleDNS":  {8, 8, 8, 8},
		"google1DNS": {4, 4, 4, 4},
	}
	for n, a := range addrs {
		fmt.Printf("%v: %v\n", n, a)
	}
}
