/*
接口
接口类型是由一组方法定义的集合。

接口类型的值可以存放实现这些方法的任何值。

注意： 列子代码的 22 行存在一个错误。 由于 Abs 只定义在 *Vertex（指针类型）上， 所以 Vertex（值类型）不满足 Abser。
*/

package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := myfloat(-math.Sqrt2)
	v := Vertex{3, 4}
	a = f
	a = &v
	fmt.Println(a.Abs())
}

type myfloat float64

func (f myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	x, y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}
