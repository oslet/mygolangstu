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
