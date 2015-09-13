package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func (v *Vertex) scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func (v *Vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := &Vertex{3, 4}
	v.scale(5)
	fmt.Println(v, v.abs())
}
